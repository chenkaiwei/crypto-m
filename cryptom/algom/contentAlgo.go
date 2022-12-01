package algom

import (
	"github.com/wumansgy/goEncrypt/aes"
	"github.com/wumansgy/goEncrypt/des"
)

type ContentAlgo interface {
	Encrypt(data []byte, cek []byte) (string, error)
	Decrypt(s string, cek []byte) ([]byte, error)
}

type contentAlgoAesCbcBase64 struct {
	ivAes []byte
}

func (ca *contentAlgoAesCbcBase64) Encrypt(data []byte, cek []byte) (string, error) {
	return aes.AesCbcEncryptBase64(data, cek, ca.ivAes)
}

func (ca *contentAlgoAesCbcBase64) Decrypt(s string, cek []byte) ([]byte, error) {
	return aes.AesCbcDecryptByBase64(s, cek, ca.ivAes)
}

/*对于AES来说PKCS5Padding和PKCS7Padding是完全一样的，不同在于PKCS5限定了块大小为8bytes而PKCS7没有限定。因此对于AES来说两者完全相同*/

type contentAlgoAesCbcHex struct {
	ivAes []byte
}

func (ca *contentAlgoAesCbcHex) Encrypt(data []byte, cek []byte) (string, error) {
	return aes.AesCbcEncryptHex(data, cek, ca.ivAes)
}

func (ca *contentAlgoAesCbcHex) Decrypt(s string, cek []byte) ([]byte, error) {
	return aes.AesCbcDecryptByHex(s, cek, ca.ivAes)
}

type contentAlgoTripleDesBase64 struct {
	ivAes []byte
}

func (ca *contentAlgoTripleDesBase64) Encrypt(data []byte, cek []byte) (string, error) {
	return des.TripleDesEncryptBase64(data, cek, ca.ivAes)
}

func (ca *contentAlgoTripleDesBase64) Decrypt(s string, cek []byte) ([]byte, error) {
	return des.TripleDesDecryptByBase64(s, cek, ca.ivAes)
}

type contentAlgoTripleDesHex struct {
	ivAes []byte
}

func (ca *contentAlgoTripleDesHex) Encrypt(data []byte, cek []byte) (string, error) {
	return des.TripleDesEncryptHex(data, cek, ca.ivAes)
}

func (ca *contentAlgoTripleDesHex) Decrypt(s string, cek []byte) ([]byte, error) {
	return des.TripleDesDecryptByHex(s, cek, ca.ivAes)
}

type contentAlgoAesCtrBase64 struct {
	ivAes []byte
}

func (ca *contentAlgoAesCtrBase64) Encrypt(data []byte, cek []byte) (string, error) {
	return aes.AesCtrEncryptBase64(data, cek, ca.ivAes)
}

func (ca *contentAlgoAesCtrBase64) Decrypt(s string, cek []byte) ([]byte, error) {
	return aes.AesCtrDecryptByBase64(s, cek, ca.ivAes)
}

type contentAlgoAesCtrHex struct {
	ivAes []byte
}

func (ca *contentAlgoAesCtrHex) Encrypt(data []byte, cek []byte) (string, error) {
	return aes.AesCtrEncryptHex(data, cek, ca.ivAes)
}

func (ca *contentAlgoAesCtrHex) Decrypt(s string, cek []byte) ([]byte, error) {
	return aes.AesCtrDecryptByHex(s, cek, ca.ivAes)
}
