package pointer

import (
	"fmt"
)

// GetPointedValue - 获取指针指向的值
func GetPointedValue() {
	fmt.Printf("\n")
	fmt.Println("++++++ 获取指针指向的值 +++++++")
	// 声明并初始化一个字符串
	var house = "Malibu Point 10880, 90265"

	// 获取 house 的内存地址
	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("ptr address: %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)

	// 	取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

	// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	// 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	// 指针变量的值是指针地址。
	// 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
}
