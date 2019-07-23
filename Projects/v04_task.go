package main

import (
	"fmt"
)

const (
	a = "A"
	b        // 使用上行的表达式，值为"A"
	c = iota // 第三个const，所以值为2
	d        // 值为3
)

const (
	e = iota // iota遇到const，重置为0
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
