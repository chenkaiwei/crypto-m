package client_demo

import (
	"bytes"
	"fmt"
	"github.com/wumansgy/goEncrypt/aes"
	"github.com/wumansgy/goEncrypt/rsa"
	"io/ioutil"
	"net/http"
	"testing"
)

//因为javascript实在不太熟，所以simpleDemo以外的客户端仅使用go做演示，不再写对应的postman用例
func TestSimpleDemo(t *testing.T) {

	aesIv := []byte("1111222233334444")

	cek := newRandomCek(16)
	fmt.Println("新生成随机cek--", cek)

	//==请求服务器阶段：
	client := &http.Client{}
	url := "http://localhost:11111/cryptionTest"

	//组装消息体
	s := "{\"msg\":\"来自TestSimpleDemo的加密测试数据\"}"
	eContent, err := aes.AesCbcEncryptHex([]byte(s), []byte(cek), aesIv)
	if err != nil {
		return
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(eContent)))
	if err != nil {
		t.Fatal(err)
	}
	//组装消息头
	eCek, err := rsa.RsaEncryptToBase64([]byte(cek), PUB_KEY)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("ECEK", eCek)

	request.Header.Set("Content-Type", "application/json")

	response, _ := client.Do(request)

	//处理响应阶段：
	defer response.Body.Close() //在回复后必须关闭回复的主体

	status := response.StatusCode
	fmt.Printf("响应状态码--%v\n", status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("响应消息体--", string(body))
	//解密响应消息体（服务端采用ResponseHandle中间件策略时）
	decryptedRespMsg, err := aes.AesCbcDecryptByHex(string(body), []byte(cek), aesIv)
	if err != nil {
		return
	}
	fmt.Println("响应消息体（解密后）--", string(decryptedRespMsg))

}
