package mydemo

import (
	"fmt"
	"testing"
)

type myF float64

func TestPlusAny(t *testing.T) {

	fmt.Println(PlusAny(3, 4))
	fmt.Println(PlusAny("3", "4"))
	fmt.Println(PlusAny("3", 4))
	fmt.Println(PlusAny("3.1", 4444.1))
	fmt.Println(PlusAny("3.1", uint16(4)))
	fmt.Println(PlusAny(myF(3333), float32(4)))

	fmt.Println(PlusAny("ss", 4))
	fmt.Println(PlusAny([]int{3}, 4))
}

func ExamplePlusAny() { //只适合单条演示，多条时相当难用，Output只能一大堆堆在后面
	sum, _ := PlusAny("3.1", uint16(4))
	fmt.Println(sum)
	plusAny, _ := PlusAny(myF(3333), float32(4))
	fmt.Println(plusAny)
	fmt.Println(PlusAny("ss", 4))

	//Output:
	//7.1
	//3337
	//<nil> 字符串不合法--无法转成数字: strconv.ParseFloat: parsing "ss": invalid syntax
}
