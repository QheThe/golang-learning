package main

import (
	"fmt"
	"golang-learning/concurrent"
	"golang-learning/device"
	"golang-learning/pointer"
	"golang-learning/loop"
	"golang-learning/sin"
	"golang-learning/strOp"
)

func main() {
	sin.Sin()

	fmt.Println("循环 ==========================================")
	loop.For_loop()

	fmt.Println("字符串操作 ==========================================")
	strOp.HasPreAndSuff()
	strOp.HasContain()

	fmt.Println("指针 ==========================================")
	pointer.MemAddress()
	pointer.GetPointedValue()
	pointer.Swap()
	pointer.GetCmdArgs()
	pointer.StructPointer()

	fmt.Println("接口 ==========================================")
	device.PhoneCall()

	fmt.Println("并发 ===========================================")
	concurrent.Ioas()
}
