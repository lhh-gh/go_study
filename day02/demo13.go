package main

import "fmt"

// 定义一个手机结构体，在Go语言中视为类，类私有成员包括：品牌、型号、颜色、价格所组成
type Phone struct {
	brand string // 品牌
	model string // 型号
	color string // 颜色
	price int    // 价格
}

// 定义一个UseingPhone 正在用的手机类，通过组合的方式集成于Phone，并有自己的类成员为
// ChargeSpeed来表示充电速度。
type UseingPhone struct {
	Phone
	ChargeSpeed string // 定义充电速度
}

// 构造方法
func NewPhone(brand, model, color string, price int) *Phone {
	return &Phone{
		brand: brand,
		model: model,
		color: color,
		price: price,
	}
}

// 封装方法获取phone的信息
func (p *Phone) GetPhoneInfo() {
	fmt.Printf("手机的品牌是：%s, 型号是：%s, 颜色是：%s, 价格是：%d\n", p.brand,
		p.model, p.color, p.price)
}

// 封装方法 获取正在使用的手机信息
func (up *UseingPhone) GetUsingPhoneInfo() {
	fmt.Printf("我是正在使用的手机，品牌是：%s, 型号是：%s, 颜色是：%s, 价格是：%d, 充电速 速是：%s\n", up.brand,
		up.model, up.color, up.price, up.ChargeSpeed)

}
func main() {
	//var phone = NewPhone("小米", "小米13 pro", "黑色", 1999)
	//phone.GetPhoneInfo()

	// 继承
	var phone = NewPhone("小米", "小米10 pro", "黑色", 3999)
	var usingPhone = UseingPhone{
		Phone:       *phone,
		ChargeSpeed: "40min",
	}
	usingPhone.GetUsingPhoneInfo()
}
