package main

import (
	"fmt"
)

func main() {
	// a := true
	// if a, b, c := 1, 2, 3; a+b+c > 6 {
	// 	fmt.Println(">6")
	// } else {
	// 	fmt.Println("<=6")
	// 	fmt.Println(a)
	// }
	// fmt.Println(a)

	// a := 1
	// 第一种，无限循环
	// for {
	// 	a++
	// 	if a > 5 {
	// 		break
	// 	}
	// 	fmt.Println(a)
	// }
	// fmt.Println("over")

	// 第二种，条件循环
	// for a < 3 {
	// 	a++
	// 	fmt.Println(a)
	// }
	// fmt.Println("over")

	// 第三种，经典
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Println("over")

	// goto的使用

	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				goto label
			}
			fmt.Println(j)
		}
	}
label:
	fmt.Println("label")
}
