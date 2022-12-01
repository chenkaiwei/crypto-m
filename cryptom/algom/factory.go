package algom

func NewCekAlgoRsaBase64(rsaPrivateKey string) CekAlgo {
	ca := &CekAlgoRsaBase64{privateKey: rsaPrivateKey}
	return ca
}
func NewCekAlgoRsaHex(rsaPrivateKey string) CekAlgo {
	ca := &CekAlgoRsaHex{privateKey: rsaPrivateKey}
	return ca
}

//====⬇️对称部分

func NewContentAlgoAesCbcBase64(ivAes []byte) ContentAlgo {
	ca := &contentAlgoAesCbcBase64{ivAes: ivAes}
	return ca
}

func NewContentAlgoAesCbcHex(ivAes []byte) ContentAlgo {
	ca := &contentAlgoAesCbcHex{ivAes: ivAes}
	return ca
}

//---

func NewContentAlgoTripleDesBase64(ivAes []byte) ContentAlgo {
	ca := &contentAlgoTripleDesBase64{ivAes: ivAes}
	return ca
}
func NewContentAlgoTripleDesHex(ivAes []byte) ContentAlgo {
	ca := &contentAlgoTripleDesHex{ivAes: ivAes}
	return ca
}

//---

func NewContentAlgoAesCtrBase64(ivAes []byte) ContentAlgo {
	ca := &contentAlgoAesCtrBase64{ivAes: ivAes}
	return ca
}
func NewContentAlgoAesCtrHex(ivAes []byte) ContentAlgo {
	ca := &contentAlgoAesCtrHex{ivAes: ivAes}
	return ca
}
