package algom

import "github.com/wumansgy/goEncrypt/rsa"

type CekAlgo interface {
	cekAlgoPrivate() //防止被其他接口无意实现
	//cek只需要解密，加密是客户端的事
	Decrypt(s string) ([]byte, error)
}

type IAmCekAlgo struct {
}

func (ca *IAmCekAlgo) cekAlgoPrivate() {
}

type CekAlgoRsaBase64 struct {
	IAmCekAlgo
	privateKey string
}

func (ca *CekAlgoRsaBase64) Decrypt(s string) ([]byte, error) {
	return rsa.RsaDecryptByBase64(s, ca.privateKey)
}

type CekAlgoRsaHex struct {
	IAmCekAlgo
	privateKey string
}

func (ca *CekAlgoRsaHex) Decrypt(s string) ([]byte, error) {
	return rsa.RsaDecryptByHex(s, ca.privateKey)
}
