package pointer

import (
	"fmt"
)

// Swap 交换变量位置
func Swap() {
	fmt.Printf("\n")
	fmt.Println("+++++++ 交换变量位置 +++++++")
	x, y := 1, 2
	swapVariabe(&x, &y)
	fmt.Println(x, y)
}

func swapVariabe(a, b *int) {
	// 取a指针的值 赋给临时变量 t 缓存
	t := *a

	// 取b指针的值 赋给a指针指向的变量
	*a = *b

	// 将a指针的值赋给b指针指向的变量
	*b = t
}
