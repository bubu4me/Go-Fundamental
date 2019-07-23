package main

import (
	"fmt"
)

// type A struct {
// 	Name string
// }

// type B struct {
// 	Name string
// }

// 定义方法：通过显示说明receiver来实现与某个类型的组合
// 此处的receiver：(a A)
// func (a A) Print() {
// 	a.Name = "aa"
// 	fmt.Println("A", a.Name)
// }

// // 此处的receiver：(b *B)
// func (b *B) Print() {
// 	b.Name = "bb"
// 	fmt.Println("B", b.Name)
// }

// type typeInt int

// func (a *typeInt) Print() {
// 	fmt.Println("typeInt")
// }

type C struct {
	name string
}

func (c *C) Print() {
	c.name = "123"
	fmt.Println(c.name)
}

func main() {
	// a := A{}
	// a.Print()                   // A aa
	// fmt.Println("Name", a.Name) // Name

	// b := B{}
	// b.Print()                   // B bb
	// fmt.Println("Name", b.Name) // Name bb

	// var a typeInt
	// a.Print()            // typeInt
	// (*typeInt).Print(&a) // typeInt

	c := C{}
	c.Print()           // 123
	fmt.Println(c.name) // 123

}
