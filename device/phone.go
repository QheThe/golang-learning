package device

import (
	"fmt"
)

// Phone 定义接口 (设备交互能力)
type Phone interface {
	call()
}

// NokiaPhone 定义结构体 (设备类)
type NokiaPhone struct{}

// 实现接口方法 (这个设备拥有 接口定义的交互能力)
// 这个方法归 NokiaPhone 这个设备
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

// PhoneCall 手机呼叫方法
func PhoneCall() {
	var nokia Phone
	// 实例化 结构体 (设备实例)
	nokia = new(NokiaPhone)
	// 调用设备交互方法
	nokia.call()
}
