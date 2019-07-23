package main

import (
	"fmt"
)

// 声明接口——一个或多个方法签名的集合
type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

// struct类型拥有所有方法签名，即实现该接口
func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func Disconnect1(usb USB) {
	// 此处还不能判断是哪个设备Disconnected
	// fmt.Println("Disconnected1")

	// 可以通过判断usb的结构来确定
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected1:", pc.name)
		return
	}
	fmt.Println("Unknown device.")
}

// 形参为空接口，即可以传递任何类型
func Disconnect2(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected2:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

func main() {
	// 定义接口变量
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect() // Connect: PhoneConnecter

	b := PhoneConnecter{"PhoneConnecter"}
	fmt.Printf("b type:%T\n", b) // b type:main.PhoneConnecter
	// 判断是否接口类型
	// Disconnect1(b) // Disconnected1 // 说明是USB类型
	// 判断实现接口的变量的类型
	Disconnect1(b) // Disconnected1: PhoneConnecter
	Disconnect2(b) // Disconnected2: PhoneConnecter

	// 接口的向下转换
	var c Connecter
	c = Connecter(a)
	c.Connect() // Connect: PhoneConnecter

	var d interface{}     // 空接口：其存储的类型和对象都为nil
	fmt.Println(d == nil) // true
	var p *int = nil      // int类型的指针，指向的对象为nil
	d = p                 // d指向p，即指向的类型为int的指针
	fmt.Println(d == nil) // false
}
