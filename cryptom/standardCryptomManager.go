package cryptom

import (
	"bytes"
	"context"
	"github.com/chenkaiwei/crypto-m/cryptom/algom"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io/ioutil"
	"net/http"
)

type standardCryptomManager struct {
	cekAlgo     algom.CekAlgo
	contentAlgo algom.ContentAlgo
}

//保证在多handle叠加使用时也仅初始化一次：若已经解密过，从context中获取；若未解密过，从header中取出cryption并解密，并将结果（内容密钥）存入context。
func (m *standardCryptomManager) getCekFromCtxOrHeader(r *http.Request) (cek []byte, reqWithCek *http.Request, err error) {
	// 更精益求精的改造还可以考虑这个方法仅用于首次解密cek时存入context，handler中后续的解密改成调用ContentDecrypt
	cek, err = getCekFromContext(r.Context())

	//⬇️CEK不存在时，从header中取出cryption并解密一遍，再存入context
	if err != nil && errors.Is(err, ErrCEKNotFoundInContext) {

		ecek := r.Header.Get("ECEK")
		logx.Info("解密内容密钥(ecek)--", ecek)
		if len(ecek) == 0 {
			err = ErrECEKNotFoundInHeader
			return
		}

		cek, err = m.cekAlgo.Decrypt(ecek)
		if err != nil {
			err = NewCryptomError(ErrTypeECEKDecryptFailure, err, "CEK解密失败")
			return
		}

		err = nil //清空err，否则会被return带出去

		logx.Info("内容密钥(cek)解密成功--", string(cek))

		reqWithCek = r.WithContext(context.WithValue(r.Context(), ContentEncryptionKey, cek)) //解好存一份，用以保证三个模块叠加使用时不会重复解密
		return
	}
	//值存在时，req不变
	reqWithCek = r

	return
}

/*custom util*/
//内容解密，手动使用
func (m *standardCryptomManager) ContentDecrypt(ctx context.Context, encryptedMsg string) (res string, err error) {

	cek, err := getCekFromContext(ctx)
	if err != nil {
		return "", err
	}

	//decryptData, err := aes.AesCbcDecryptByHex(encryptedMsg, tempKey, m.aesIv)
	decryptData, err := m.contentAlgo.Decrypt(encryptedMsg, cek)
	if err != nil {

		return "", NewCryptomError(ErrTypeContentDecryptFailure, err, "消息正文(content)解密出错")
	}
	res = string(decryptData)
	return
}

//内容加密
func (m *standardCryptomManager) ContentEncrypt(ctx context.Context, msg string) (res string, err error) {

	cek, err := getCekFromContext(ctx)
	if err != nil {
		return
	}

	//encryptedData, err := aes.AesCbcEncryptHex([]byte(msg), cek, m.aesIv)
	res, err = m.contentAlgo.Encrypt([]byte(msg), cek)
	if err != nil {
		return "", NewCryptomError(ErrTypeContentEncryptFailure, err, "消息正文(content)加密出错")
	}
	return
}

func (m *standardCryptomManager) RequestHandle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		logx.Info(" CryptomManager RequestHandle")
		cek, reqWithCek, err := m.getCekFromCtxOrHeader(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		//解析全body
		body, err := ioutil.ReadAll(reqWithCek.Body)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		logx.Info("原始消息体--", string(body))
		//解密
		decryptBody, err := m.contentAlgo.Decrypt(string(body), cek)
		if err != nil {
			//httpx.Error(w, errors.WithStack(NewCryptomError(ErrTypeContentDecryptFailure, err, "正文解密出错")))//要不要带调用栈？
			httpx.Error(w, NewCryptomError(ErrTypeContentDecryptFailure, err, "正文解密出错"))
			return
		}
		logx.Info("解密后的消息体--", decryptBody)
		closer := ioutil.NopCloser(bytes.NewBuffer(decryptBody))
		//写回body
		reqWithCek.Body = closer

		next(w, reqWithCek)
	}
}

//仅将Cek解密后存入context，在logic中通过manager手动解密加密
func (m *standardCryptomManager) ManualHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, newReq, err := m.getCekFromCtxOrHeader(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		// Passthrough to next handler if need
		next(w, newReq)
	}
}

//对整个返回的响应body进行加密
func (m *standardCryptomManager) ResponseHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cek, reqWithCek, err := m.getCekFromCtxOrHeader(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		logx.Info("cek--", cek)

		writerDataReciever := newResponseWriterDataReciever(w)
		//
		next(writerDataReciever, reqWithCek) //传newReq，令仅应用ResponseHandle时也能有CustomHandle的效果

		header := writerDataReciever.Header()
		logx.Info("header--", header)

		logx.Info("加密前的响应消息--" + string(writerDataReciever.Data))

		//⬇️ 错误（状态码不为200）时不加密，直接返回（业务逻辑错误时也是200「gozero官方示例的做法」，所以会和正常状态一样加密。如果需要改成不加密，在httpx.SetErrorHandler中国呢另行定义一个statusCode即可）
		//statusCode！=200时直接返回
		if writerDataReciever.StatusCode != http.StatusOK {
			w.Write(writerDataReciever.Data)
			return
		}

		//加密整个返回消息体
		encryptedData, err := m.contentAlgo.Encrypt(writerDataReciever.Data, cek)

		if err != nil {
			httpx.Error(w, errors.WithStack(NewCryptomError(ErrTypeContentEncryptFailure, err, "返回值加密发生错误")))
			return
		}

		logx.Info("加密后的响应消息---", encryptedData)

		w.Write([]byte(encryptedData))

	}
}

///*@responseWriterDataReciever 内部类部分*/

//因为ResponseWriter不能直接访问，所以继承一下，接收时存到公共属性
//write方法不在事实上执行写操作，仅代替writer接收数据，加密后再执行写入
//httptest.ResponseRecorder 也有类似功能，虽然拿来用也行，但那个好像是模拟请求用的，还是自己写个吧更为可控）
//
type responseWriterDataReciever struct {
	http.ResponseWriter
	Data       []byte
	StatusCode int
}

func newResponseWriterDataReciever(w http.ResponseWriter) *responseWriterDataReciever {
	return &responseWriterDataReciever{
		w, nil, 200,
	}
}

//不在事实上执行写操作，仅代替writer接收数据，加密后再执行写入
func (w *responseWriterDataReciever) Write(i []byte) (res int, err error) {
	w.Data = i
	res = len(w.Data)
	return
}

func (w *responseWriterDataReciever) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterDataReciever) GetData() []byte {
	return w.Data
}

//responseWriterDataReciever 内部类部分 end
