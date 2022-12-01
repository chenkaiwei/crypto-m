package cryptomx

import (
	"github.com/pkg/errors"
)

type MyCaesarAlgo struct {
}

/*
本演示用的自定义加密策略，古老的凯撒加密法：
cek为偏移量，加密过程为每个字节byte在原数字的基础上加上cek，超出2^8时循环回1计数
*/

func (ca *MyCaesarAlgo) Encrypt(data []byte, cek []byte) (res string, err error) {
	offset := int(cek[0])
	if err != nil || offset < 0 || offset > 255 {
		return "", errors.WithMessage(err, "cek不合法，请确保其为一个0~255范围内的数字")
	}

	for i, datum := range data {
		sum := int(datum) + offset
		if sum >= 256 {
			sum = sum - 256
		}
		data[i] = byte(sum)
	}
	res = string(data)
	return
}

func (ca *MyCaesarAlgo) Decrypt(s string, cek []byte) (res []byte, err error) {

	offset := int(cek[0])
	if err != nil || offset < 0 || offset > 255 {
		return nil, errors.WithMessage(err, "cek不合法，请确保其为一个0~255范围内的数字")
	}
	data := []byte(s)
	for i, datum := range data {
		sum := int(datum) - offset
		if sum < 0 {
			sum = sum + 256
		}
		data[i] = byte(sum)
	}
	res = data
	return
}
