package main

import "fmt"

// 定义手机功能
type Phones interface {
	Call(to string) // 打电话
	OnInternet()    // 上网
}

// 定义苹果手机结构体
type IPhone struct{}

// 苹果手机实现打电话功能
func (p *IPhone) Call(to string) {
	fmt.Printf("使用IPhone手机打电话给：%s\n", to)
}

// 苹果手机实现上网功能
func (p *IPhone) OnInternet() {
	fmt.Println("使用IPhone手机上网")
}

// 定义安卓手机结构体
type Android struct{}

// 安卓手机实现打电话功能
func (p *Android) Call(to string) {
	fmt.Printf("使用Android手机打电话给：%s\n", to)
}

// 安卓手机实现上网功能
func (p *Android) OnInternet() {
	fmt.Println("使用Android手机上网")
}

// 打电话
func PCall(p Phones, to string) {
	p.Call(to)
}

// 上网
func POnInternet(p Phones) {
	p.OnInternet()
}
func main() {
	// 多态
	var iphone = &IPhone{}
	PCall(iphone, "包子")
	POnInternet(iphone)
	fmt.Println()
	var android = &Android{}
	PCall(android, "肉包子")
	POnInternet(android)
}
