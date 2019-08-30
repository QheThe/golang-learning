package pointer

import (
	"flag"
	"fmt"
)

// GetCmdArgs - 获取命令行参数
func GetCmdArgs() {
	fmt.Printf("\n")
	fmt.Println("+++++++ 获取命令行参数 ++++++++")
	// 定义命令行参数
	var mode = flag.String("mode", "", "process mode")

	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数
	fmt.Println(*mode)
}
