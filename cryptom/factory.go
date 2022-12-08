package cryptom

import "github.com/chenkaiwei/crypto-m/cryptom/algom"

//工厂模式 别的语言爱用Create，但go爱用New，入乡随俗吧
func NewDefaultCryptomManager(rsaPrivateKey string, ivAes []byte) CryptomManager {

	//旧 写死的初始化
	//return &DefaultCryptionManager{
	//	rsaPrivateKey: rsaPrivateKey,
	//	//↑ RSA-base64，用于解密内容密钥
	//	aesIv: ivAes,
	//	//↑ 用于解正文
	//}

	//新 间接调用standard的初始化，废弃原DefaultCryptionManager.go
	cekAlgo := algom.NewCekAlgoRsaBase64(rsaPrivateKey)
	contentAlgo := algom.NewContentAlgoAesCbcHex(ivAes)
	cryptomManager := NewStandardCryptomManager(cekAlgo, contentAlgo)
	return cryptomManager
}

func NewStandardCryptomManager(cekAlgo algom.CekAlgo, contentAlgo algom.ContentAlgo) CryptomManager {
	return &standardCryptomManager{
		cekAlgo, contentAlgo,
	}
}
