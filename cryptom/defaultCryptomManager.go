package cryptom

// 本文件已被standardCryptomManager取代，源码缓缓再删
//
//type DefaultCryptionManager struct {
//	rsaPrivateKey string
//	aesIv         []byte
//}
//
////保证在多handle叠加使用时也仅初始化一次：若已经解密过，从context中获取；若未解密过，从header中取出cryption并解密，并将结果（内容密钥）存入context。
//func (m *DefaultCryptionManager) getCekFromCtxOrHeader(r *http.Request) (cek []byte, reqWithTempKey *http.Request, err error) {
//
//	cek, err = getCekFromContext(r.Context())
//
//	//⬇️CEK不存在时，从header中取出cryption并解密一遍，再存入context
//	//var cryptomError *CryptomError
//	//if err != nil && errors.As(err, &cryptomError) && cryptomError == ErrCEKNotFoundInContext {
//	if err != nil && errors.Is(err, ErrCEKNotFoundInContext) {
//
//		cryption := r.Header.Get("ECEK")
//		logx.Info("解密内容密钥--", cryption)
//		if len(cryption) == 0 {
//			//err = errorx.NewCodeError(errorx.CRYPTION_NOT_FOND, "header中CRYPTION字段不存在")
//			err = ErrECEKNotFoundInHeader
//			return
//		}
//
//		cek, err = rsa.RsaDecryptByBase64(cryption, m.rsaPrivateKey)
//		if err != nil {
//			//err = errorx.NewCodeError(errorx.TEMP_KEY_DECRYPT_ERROR, "内容密钥解密错误")
//			err = ErrECEKDecryptFailure
//			return
//		}
//
//		err = nil //清空err，否则会被return带出去
//
//		logx.Info("内容密钥解密成功--", string(cek))
//
//		reqWithTempKey = r.WithContext(context.WithValue(r.Context(), ContentEncryptionKey, cek)) //解好存一份，用以保证三个模块叠加使用时不会重复解密
//		return
//	}
//	//值存在时，req不变
//	reqWithTempKey = r
//
//	return
//}
//
///*custom util*/
////单行解密
//func (m *DefaultCryptionManager) ContentDecrypt(ctx context.Context, encryptedMsg string) (res string, err error) {
//
//	tempKey, err := getCekFromContext(ctx)
//	if err != nil {
//		return
//	}
//
//	decryptData, err := aes.AesCbcDecryptByHex(encryptedMsg, tempKey, m.aesIv)
//	if err != nil {
//		err = NewCryptomError(ErrTypeContentDecryptFailure, err, "消息体解密失败")
//		return
//	}
//	res = string(decryptData)
//	return
//}
//
////单行加密
//func (m *DefaultCryptionManager) ContentEncrypt(ctx context.Context, msg string) (res string, err error) {
//
//	tempKey, err := getCekFromContext(ctx)
//	if err != nil {
//		return
//	}
//
//	encryptedData, err := aes.AesCbcEncryptHex([]byte(msg), tempKey, m.aesIv)
//	if err != nil {
//		return
//	}
//
//	res = string(encryptedData)
//	return
//}
//
//func (m *DefaultCryptionManager) RequestHandle(next http.HandlerFunc) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		logx.Info(" CryptomManager RequestHandle")
//		tempAesCbcKey, reqWithTempKey, err := m.getCekFromCtxOrHeader(r)
//		if err != nil {
//			httpx.Error(w, err)
//			return
//		}
//
//		//解析全body
//		body, err := ioutil.ReadAll(reqWithTempKey.Body)
//		if err != nil {
//
//			httpx.Error(w, err)
//			return
//		}
//		logx.Info("原始消息体--", string(body))
//		//解密
//		decryptBody, err := aes.AesCbcDecryptByHex(string(body), tempAesCbcKey, m.aesIv)
//		//decryptBody, err := aes.AesCbcDecryptByBase64(string(body), tempAesCbcKey, m.aesIv)
//		if err != nil {
//			err = NewCryptomError(ErrTypeContentDecryptFailure, err, "消息体解密失败")
//			httpx.Error(w, err)
//			return
//		}
//		logx.Info("解密后的消息体--", decryptBody)
//
//		/*对于AES来说PKCS5Padding和PKCS7Padding是完全一样的，不同在于PKCS5限定了块大小为8bytes而PKCS7没有限定。因此对于AES来说两者完全相同*/
//
//		logx.Info("decryptedBodyStr--", string(decryptBody))
//		closer := ioutil.NopCloser(bytes.NewBuffer(decryptBody))
//		//写回body
//		reqWithTempKey.Body = closer
//
//		next(w, reqWithTempKey)
//	}
//}
//
////仅将临时密钥解密后存入context，供后续在logic中使用（通过GetTempAesCbcKeyFromContext方法取出）
//func (m *DefaultCryptionManager) ManualHandle(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		_, newReq, err := m.getCekFromCtxOrHeader(r)
//		if err != nil {
//			fmt.Fprintln(w, err.Error()) //TOOD 后期改进为和errorx格式一致
//			return
//		}
//		// Passthrough to next handler if need
//		next(w, newReq)
//	}
//}
//
////对整个返回的响应body进行加密
//func (m *DefaultCryptionManager) ResponseHandle(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		tempAesCbcKey, reqWithTempKey, err := m.getCekFromCtxOrHeader(r)
//		if err != nil {
//			//fmt.Fprintln(w, err.Error()) //中间件中返回错误
//			httpx.Error(w, err)
//			return
//		}
//		logx.Info("tempAesCbcKey--", tempAesCbcKey)
//
//		writerDataReciever := newResponseWriterDataReciever(w)
//		//
//		next(writerDataReciever, reqWithTempKey) //传newReq，令仅应用ResponseHandle时也能有CustomHandle的效果
//
//		header := writerDataReciever.Header()
//		logx.Info("header--", header)
//
//		logx.Info("加密前的响应消息--" + string(writerDataReciever.Data))
//
//		//⬇️ 错误（状态码不为200）时不加密，直接返回（业务逻辑错误时也是200「gozero官方示例的做法」，所以会和正常状态一样加密。如果需要改成不加密，在httpx.SetErrorHandler中国呢另行定义一个statusCode即可）
//		//statusCode！=200时直接返回
//		if writerDataReciever.StatusCode != http.StatusOK {
//			w.Write(writerDataReciever.Data)
//			return
//		}
//
//		//加密整个返回消息体
//		//datatest := []byte(string(writerDataReciever.Data) + "ddd") //测试 非json格式的报错
//		//encryptedData, err := aes.AesCbcEncryptHex(datatest, tempAesCbcKey, m.aesIv)
//
//		encryptedData, err := aes.AesCbcEncryptHex(writerDataReciever.Data, tempAesCbcKey, m.aesIv)
//		//encryptedData, err := aes.AesCbcEncryptBase64(writerDataReciever.Data, tempAesCbcKey, m.aesIv)
//
//		if err != nil {
//			httpx.Error(w, errors.WithMessage(err, "返回值加密发生错误"))
//			return
//		}
//
//		logx.Info("加密后的响应消息---", encryptedData)
//
//		w.Write([]byte(encryptedData))
//
//	}
//}
