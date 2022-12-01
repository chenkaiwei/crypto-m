package client_demo

import (
	"bytes"
	"fmt"
	"github.com/wumansgy/goEncrypt/rsa"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCustomDemo(t *testing.T) {

	cek := 5

	//==请求服务器阶段：
	client := &http.Client{}
	url := "http://localhost:11113/customDemo"

	//组装消息体
	s := "{\"msg\":\"来自TestCustomDemo的测试消息\"}"

	//加密
	offset := cek
	data := []byte(s)
	for i, datum := range data {
		sum := int(datum) + offset
		if sum >= 256 {
			sum = sum - 256
		}
		data[i] = byte(sum)
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		t.Fatal(err)
	}
	//组装消息头
	eCek, err := rsa.RsaEncryptToBase64([]byte{byte(cek)}, PUB_KEY)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("ECEK", eCek)

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("响应返回")
	//处理响应阶段：
	defer response.Body.Close() //在回复后必须关闭回复的主体

	status := response.StatusCode
	fmt.Printf("响应状态码--%v\n", status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("响应消息体--", string(body))

	for i, datum := range body {
		sum := int(datum) - offset
		if sum < 0 {
			sum = sum + 256
		}
		body[i] = byte(sum)
	}
	fmt.Println("解密后的响应消息体--", string(body))

}
