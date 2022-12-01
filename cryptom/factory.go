package cryptom

import "github.com/chenkaiwei/crypto-m/cryptom/algom"

func CreateDefaultCryptomManager(rsaPrivateKey string, ivAes []byte) CryptomManager {

	//旧 写死的初始化
	//return &DefaultCryptionManager{
	//	rsaPrivateKey: rsaPrivateKey,
	//	//↑ RSA-base64，用于解密内容密钥
	//	aesIv: ivAes,
	//	//↑ 用于解正文
	//}

	//新 间接调用standard的初始化，废弃原DefaultCryptionManager
	cekAlgo := algom.NewCekAlgoRsaBase64(rsaPrivateKey)
	contentAlgo := algom.NewContentAlgoAesCbcHex(ivAes)
	cryptomManager := CreateStandardCryptomManager(cekAlgo, contentAlgo)
	return cryptomManager
}

func CreateStandardCryptomManager(cekAlgo algom.CekAlgo, contentAlgo algom.ContentAlgo) CryptomManager {
	return &StandardCryptomManager{
		cekAlgo, contentAlgo,
	}
}
