package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	A
	Name string
}

type C struct {
	Name string
}

// 当有两个及以上的结构嵌入时，它们属于同一级别，不分先后顺序
type D struct {
	A
	C
}

func main() {
	a := B{Name: "B", A: A{Name: "A"}}
	fmt.Println(a)      // {{A} B}
	fmt.Println(a.Name) // B

	d := D{A: A{Name: "A"}, C: C{Name: "C"}}
	fmt.Println(d) //{{A} {C}}
	// 由于D有两个及以上的结构嵌入，它们不分先后顺序，故d.name无法确定去找哪个嵌入结构的字段，引发错误
	// fmt.Println(d.Name)
}
