package mydemo

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

//以前写的通用加法，兼容string，用来测试包发布

func convertAnyToFloat(a any) (aNum float64, err error) {

	va := reflect.ValueOf(a)
	if va.CanFloat() {
		aNum = va.Float()
	} else if va.CanInt() {
		aNum = float64(va.Int())
	} else if va.CanUint() {
		aNum = float64(va.Uint())
	} else if as, ok := a.(string); ok {
		aNum, err = strconv.ParseFloat(as, 64)
		if err != nil {
			return 0, errors.WithMessage(err, "字符串不合法--无法转成数字")
		}
	} else {
		return 0, errors.New("不支持的类型--" + va.Type().String())
	}
	return
}

//// 在plus⬆️的基础上，令传入字符串时也转成数字后相加--尝试2，用反射
func PlusAny(a, b any) (res any, err error) {

	aNum, err := convertAnyToFloat(a)
	if err != nil {
		return nil, err
	}
	bNum, err := convertAnyToFloat(b)
	if err != nil {
		return nil, err
	}
	return aNum + bNum, nil
}
