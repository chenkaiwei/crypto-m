package client_demo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wumansgy/goEncrypt/aes"
	"github.com/wumansgy/goEncrypt/des"
	"github.com/wumansgy/goEncrypt/rsa"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMulti1Demo(t *testing.T) {

	//==请求服务器阶段：
	cek := newRandomCek(24)
	fmt.Println("随机cek--", cek)

	client := &http.Client{}
	url := "http://localhost:11112/multiDemo1"

	//组装消息体
	//	multi1 (algom.NewCekAlgoRsaBase64(PRIVATE_KEY), algom.NewContentAlgoTripleDesBase64(nil))
	s := "{\"msg\":\"来自TestMulti1Demo的加密测试数据aaaa\"}"
	eContent, err := des.TripleDesEncryptBase64([]byte(s), []byte(cek), nil)
	if err != nil {
		t.Fatal(err)
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

	response, err := client.Do(request)
	defer response.Body.Close() //在回复后必须关闭回复的主体
	//处理响应阶段：
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("响应返回")
	}

	status := response.StatusCode
	fmt.Printf("状态码--%v\n", status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("响应消息体--", string(body))
	}
	//解密响应消息体（服务端采用ResponseHandle中间件策略时）
	//decryptedRespMsg, err := aes.AesCbcDecryptByHex(string(body), []byte(cek), nil)
	//des.TripleDesEncryptBase64([]byte(s), []byte(cek), nil)
	decryptedRespMsg, err := des.TripleDesDecryptByBase64(string(body), []byte(cek), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("响应消息体（解密后）--", string(decryptedRespMsg))

}

func TestMulti2Demo(t *testing.T) {

	//==请求服务器阶段：
	cek := newRandomCek(24)
	fmt.Println("随机cek--", cek)

	client := &http.Client{}
	url := "http://localhost:11112/multiDemo2"

	//组装消息体
	//	multi2CryptomManager (algom.NewCekAlgoRsaBase64(PRIVATE_KEY), algom.NewContentAlgoAesCtrHex(nil)
	s := "{\"msg\":\"来自TestMulti2Demo的加密测试数据bbbb\"}"

	eContent, err := aes.AesCtrEncryptHex([]byte(s), []byte(cek), nil)
	if err != nil {
		t.Fatal(err)
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

	response, err := client.Do(request)
	defer response.Body.Close() //在回复后必须关闭回复的主体
	//处理响应阶段：
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("响应返回")
	}

	status := response.StatusCode
	fmt.Printf("状态码--%v\n", status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("响应消息体--", string(body))
	}
	fmt.Println("（服务端未采用Response handle，无须解密）")
}

func TestManualDemo(t *testing.T) {

	//==请求服务器阶段：
	cek := newRandomCek(24)
	fmt.Println("随机cek--", cek)

	client := &http.Client{}
	url := "http://localhost:11112/manualDemo"

	//组装消息体
	//	multi1 (algom.NewCekAlgoRsaBase64(PRIVATE_KEY), algom.NewContentAlgoTripleDesBase64(nil))
	encryptedMsg, err := des.TripleDesEncryptBase64([]byte("局部加密测试专用字符串"), []byte(cek), nil)
	reqBody := map[string]string{
		"msg":          "手动加密接口中不加密的信息",
		"encryptedMsg": encryptedMsg,
	}
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
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

	response, err := client.Do(request)
	defer response.Body.Close() //在回复后必须关闭回复的主体
	//处理响应阶段：
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("响应返回")
	}

	status := response.StatusCode
	fmt.Printf("状态码--%v\n", status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("响应消息体--", string(body))
	}

	//解密响应消息体中的加密部分（事先与服务端约定好字段如这里的encryptedMsg）
	var respMap map[string]string
	err = json.Unmarshal(body, &respMap)
	if err != nil {
		t.Fatal(err)
	}
	manualCryptoMsg, err := des.TripleDesDecryptByBase64(respMap["encryptedMsg"], []byte(cek), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("来自服务端的局部加密数据-- " + string(manualCryptoMsg))

}
