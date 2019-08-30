package pointer

import "fmt"

// MemAddress - how to get memory address of an variable
func MemAddress() {
	fmt.Printf("\n")
	fmt.Println("+++++++ 获取内存地址 +++++++")

	var car = 1
	var str = "banana"

	fmt.Printf("car: %p \n", &car)
	fmt.Printf("str: %p \n", &str)
}
