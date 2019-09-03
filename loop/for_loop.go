package loop

import (
	"fmt"
)

func while_loop() {
	var i int
	// 无条件 for 等价于 while true
	for {
		i := i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		// 条件判断
		if i < 0 {
			// 跳出循环
			break
		}
	}
}

// For_loop - for 循环
func For_loop() {
	fmt.Println("++++++ for 循环 +++++++")
	while_loop()
}