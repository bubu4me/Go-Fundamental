package main

import (
	"fmt"
)

func add1(s []int) {
	s = append(s, 3)
	fmt.Println(s)
}

func add2(s []int) []int {
	s = append(s, 3)
	return s
}

func main() {
	s := make([]int, 0)
	fmt.Println(s) // []
	add1(s)        // [3]
	fmt.Println(s) // []
	/*
		执行add()方法后没有效果
		分析：初始容量为0，经过add()的append后，
		容量增大，此时导致内存重新分配，
		add()内的s的内存地址已经发生改变，
		需要通过返回值来获取
	*/
	fmt.Println(add2(s)) // [3]
}
