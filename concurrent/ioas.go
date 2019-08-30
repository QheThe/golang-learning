package concurrent

import (
	"fmt"
	"time"
)

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}

// Ioas - 输出的同时接受输入
func Ioas() {
	fmt.Printf("\n")
	fmt.Println("+++++++++ 输出的同时接受输入 ++++++++")
	// 并发执行程序
	go running()
	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
}
