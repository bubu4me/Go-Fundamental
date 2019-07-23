package main

import (
	"fmt"
)

func main() {
	// A(10)                  // 10
	// fmt.Println(B(10))     // 10 10
	// fmt.Println(C(10, 11)) // 10 11
	// D(1, 2)                // [1 2]
	// D(1, 2, 3, 4, 5)       // [1 2 3 4 5]

	// // 值传递
	// a, b := 1, 2
	// E(1, 2)                   //函数内： [3 4]
	// fmt.Println("函数外：", a, b) // 函数外： 1 2

	// // slice的内存地址拷贝
	// s1 := []int{1, 2, 3, 4}
	// F(s1)                   // 函数内： [5 6 7 8]
	// fmt.Println("函数外：", s1) // 函数外： [5 6 7 8]

	// // 数据类型的内存地址拷贝
	// a1 := 1
	// G(&a1)
	// fmt.Println("函数外", a1) // 函数外 2

	// 匿名函数
	// f1 := func() {
	// 	fmt.Println("匿名函数f1")
	// }
	// f1() // 匿名函数f1

	// // 闭包引用外部变量
	// n := 1
	// f2 := func() int {
	// 	n++
	// 	return n
	// }
	// fmt.Println("闭包引用外部变量：", f2()) // 闭包引用外部变量： 2
	// fmt.Println(n)                 // 2

	// // 闭包作为函数返回值
	// f3 := closure(10)
	// fmt.Println(f3(1)) // 11
	// fmt.Println(f3(2)) // 12

	// // defer的调用顺序
	// fmt.Print("a ")
	// defer fmt.Print("b ")
	// defer fmt.Print("c ")
	// // a c b
	// // 可见defer调用顺序相反

	// for i := 0; i < 3; i++ {
	// 	defer fmt.Print(i, " ") // 2 1 0
	// }

	// defer匿名函数（闭包）
	// for i := 0; i < 3; i++ {
	// 	defer func() {
	// 		fmt.Print(i, " ")
	// 	}() // 加()相当于调用匿名函数
	// }

	// panic
	// pA() // Func pA
	// pB() // panic: Panic in B
	// pC()

	// panic/recover
	pD() // Func pD
	pE() // Recover in pE
	pF() // Func pF
}

// // 带参无返回值函数
// func A(a int) {
// 	fmt.Println(a)
// }

// // 带参有返回值函数——第一种形式,最好声明返回值名称
// func B(a int) (b, c int) {
// 	b, c = a, a
// 	return b, c
// }

// // 带参有返回值函数——第二种形式
// func C(a, b int) (int, int) {
// 	return a, b
// }

// // 可变长度参数,规定必须放在参数列表中的最右边位置
// func D(a ...int) {
// 	fmt.Println(a) // 可变长度参数相当于一个可变长的数组
// }

// // 值传递
// func E(a ...int) {
// 	a[0] = 3
// 	a[1] = 4
// 	fmt.Println("函数内：", a) // 函数内： [3 4]
// }

// // 内存地址拷贝
// func F(s []int) {
// 	s[0] = 5
// 	s[1] = 6
// 	s[2] = 7
// 	s[3] = 8
// 	fmt.Println("函数内：", s) // 函数内： [5 6 7 8]
// }

// // 数据类型的内存地址拷贝
// func G(a *int) {
// 	*a = 2
// 	fmt.Println("函数内", *a) // 函数内 2
// }

// 闭包
// func closure(x int) func(int) int {
// 	fmt.Printf("%p\n", &x) // 0xc0000580a8
// 	return func(y int) int {
// 		fmt.Printf("%p\n", &x) // 0xc0000580a8
// 		return x + y
// 	}
// }

// func pA() {
// 	fmt.Println("Func pA")
// }

// func pB() {
// 	panic("Panic in B")
// }

// func pC() {
// 	fmt.Println("Func C")
// }

func pD() {
	fmt.Println("Func pD")
}

func pE() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in pE")
		}
	}()

	panic("Panic in pE")
}

func pF() {
	fmt.Println("Func pF")
}
