package main

import (
	"fmt"
)

func main() {
	// // 直接创建slice
	// var s1 []int
	// fmt.Println(s1) // []

	// // 从底层数组获取slice
	// // 首先创建数组
	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(a) // [1 2 3 4 5 6 7 8 9 0]
	// // 截取数组创建slice
	// s2 := a[5:10]   // 注意索引值的含义：5代表从下表为5的元素开始截取，取几个元素呢，由(右值-左值)决定
	// fmt.Println(s2) // [6 7 8 9 0]

	// s3 := a[5:]     // 取数组的后5个元素
	// fmt.Println(s3) // [6 7 8 9 0]

	// s4 := a[:5]     // 取数组的前5个元素
	// fmt.Println(s4) // [1 2 3 4 5]

	// s := a[:] // 创建一个完整数组的切片
	// fmt.Println(s) // [1 2 3 4 5 6 7 8 9 0]

	// s5 := make([]int, 3)          // 长度len=3,未定义容量cap
	// fmt.Println(len(s5), cap(s5)) // 输出3 3，说明如果未定义容量，则cap=len
	// fmt.Println(s5)               // [0 0 0]

	// s6 := make([]int, 3, 6) // len=3,cap=10
	// fmt.Println(s6)         // [0 0 0]

	// c := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}
	// 指向同一个数组
	// sa := c[3:7]
	// fmt.Println("sa:", string(sa), "len:", len(sa), "cap:", cap(sa)) // sa: defg len: 4 cap: 9
	// sb := c[4:6]
	// fmt.Println("sb:", string(sb), "len:", len(sb), "cap:", cap(sb)) // sb: ef len: 2 cap: 8
	// fmt.Println("sa中e的地址：", &sa[1])                                  // sa中e的地址： 0xc000058084
	// fmt.Println("sb中e的地址：", &sb[0])                                  // sb中e的地址： 0xc000058084

	// 切分slice
	// sc := c[4:6]
	// fmt.Println("sc:", string(sc), "len:", len(sc), "cap:", cap(sc)) // sc: ef len: 2 cap: 8
	// sd := sc[1:5]
	// fmt.Println("sd:", string(sd), "len:", len(sd), "cap:", cap(sd)) // sd: fghi len: 4 cap: 7

	// // slice的append操作
	// se := make([]int, 3, 6)
	// fmt.Printf("%p\n", se) // 0xc0000103c0
	// se = append(se, 1, 2, 3)
	// fmt.Printf("%v %p\n", se, se) // [0 0 0 1 2 3] 0xc0000103c0
	// se = append(se, 4, 5, 6)
	// fmt.Printf("%v %p\n", se, se) // [0 0 0 1 2 3 4 5 6] 0xc00000a120

	// copy(dst,src)
	// len(s1)>len(s2)
	// s1 := []int{1, 2, 3, 4, 5, 6}
	// s2 := []int{7, 8, 9}
	// copy(s1, s2)
	// fmt.Println("s1:", s1, "s2:", s2) //s1: [7 8 9 4 5 6] s2: [7 8 9]

	// len(s1)<len(s2)
	// s1 := []int{7, 8, 9}
	// s2 := []int{1, 2, 3, 4, 5, 6}
	// copy(s1, s2)
	// fmt.Println("s1:", s1, "s2:", s2) //s1: [1 2 3] s2: [1 2 3 4 5 6]

}
