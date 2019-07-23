package main

import (
	"fmt"
)

type person1 struct {
}

type person2 struct {
	Name string
	Age  int
}

type person3 struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

type person4 struct {
	string
	int
}

type human struct {
	Sex string
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	// p1 := person1{}
	// fmt.Println(p1) // {}

	// p2 := person2{}
	// fmt.Println(p2) // { 0}

	// // 给结构体赋值
	// // 方式1
	// p2.Name = "Go"
	// p2.Age = 9
	// fmt.Println(p2) // {Go 9}
	// // 方式2
	// p3 := person2{
	// 	Name: "Go",
	// 	Age:  9,
	// }
	// fmt.Println(p3) // {Go 9}

	// p4 := person2{
	// 	Name: "Go",
	// 	Age:  9,
	// }
	// // 值传递
	// fmt.Println(p4) // {Go 9}
	// A(p4)           // {Go 11}
	// fmt.Println(p4) // {Go 9}

	// // 内存地址传递
	// // 方式1——对变量取地址
	// B(&p4)          // &{Go 11}
	// fmt.Println(p4) // {Go 11}

	// // 方式2——初始化时取地址
	// p5 := &person2{
	// 	Name: "Go",
	// 	Age:  13,
	// }
	// fmt.Println(p5) // &{Go 13}
	// B(p5)           // &{Go 11}
	// fmt.Println(p5) // &{Go 11}
	// p5.Name = "go"  // 可直接用变量名修改结构体的属性
	// fmt.Println(p5) // &{go 11}

	// 通过匿名结构声明
	// p6 := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "Go",
	// 	Age:  9,
	// }
	// fmt.Println(p6) // {Go 9}

	// 嵌套struct的赋值
	// p7 := person3{Name: "Go", Age: 15}
	// p7.Contact.Phone = "111111111"
	// p7.Contact.City = "beijing"
	// fmt.Println(p7)

	// struct的匿名字段
	// p8 := person4{"Go", 15}
	// fmt.Println(p8) // {Go 15}

	// 结构体的嵌入
	// 第一种赋值方法
	t := teacher{
		Name:  "t",
		Age:   20,
		human: human{Sex: "女"}}
	s := student{
		Name:  "s",
		Age:   10,
		human: human{Sex: "男"}}
	fmt.Println(t, s) // {{女} t 20} {{男} s 10}
	// 第二种的赋值方法
	t.human.Sex = "female"
	fmt.Println(t) // {{female} t 20}
	// 第三种的赋值方法
	s.Sex = "male"
	fmt.Println(s) // {{male} s 10}

}

// func A(per person2) {
// 	per.Age = 11
// 	fmt.Println(per)
// }

// func B(per *person2) {
// 	per.Age = 11
// 	fmt.Println(per)
// }
