package main

import (
	"fmt"
)

type A struct {
	int
}

func (a *A) Increase(num int) {
	a.int += num
}

// 优化
type TZ int

func (tz *TZ) Increase(num int) {
	*tz += TZ(num) // 虽然*tz和num底层都是int类型，但是相加时仍然要通过类型转换为相同类型
}

func main() {
	a := A{}
	a.Increase(100)
	fmt.Println(a) // {100}

	var t TZ
	t.Increase(100)
	fmt.Println(t) // 100
}
