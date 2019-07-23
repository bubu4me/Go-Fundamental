package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a)

	// 调用Itoa():数字型转字符串
	c := strconv.Itoa(a)
	// 转回来
	a, _ = strconv.Atoi(c)

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a)
}
