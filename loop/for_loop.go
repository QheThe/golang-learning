package loop

import (
	"fmt"
)

func continue_section() {
	for i:=0; i<10; i++ {
		if i == 5 {
			// 跳过当前循环
			fmt.Println("skip 5")
			continue
		}
		fmt.Print(i)
		fmt.Print("  ")
	}
}

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
	fmt.Println("")
}

// break_inner - break 语句只会跳出单层循环
func break_inner() {
	// 循环 3次 输出一次
	for i :=0; i<3; i++ {
		// 阻塞 j < 10 输出10次
		for j :=0; j<10; j++ {
			fmt.Print(j)
		}
		// 输出空格 分行
		fmt.Print("  ")
	}
	fmt.Println("")
}

// For_loop - for 循环
func For_loop() {
	fmt.Println("++++++ for 循环 +++++++")
	while_loop()
	break_inner()
	continue_section()
}