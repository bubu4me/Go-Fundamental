package main

import (
	"fmt"
	"sort"
)

func main() {
	// map的创建方法
	// var m1 map[int]string
	// m1 = map[int]string{}
	// fmt.Println(m1) // map[]

	// var m2 map[int]string
	// m2 = make(map[int]string)
	// fmt.Println(m2) // map[]

	// m3 := map[int]string{}
	// fmt.Println(m3) // map[]
	// m4 := make(map[int]string)
	// fmt.Println(m4) // map[]

	// 展示添加、删除键值对，以及获取key的对应value的操作
	// m := make(map[int]string)
	// m[1] = "hello"    // 添加key-value
	// fmt.Println(m)    // map[1:hello]
	// val1 := m[1]      // 取出key=1的value
	// fmt.Println(val1) // hello
	// delete(m, 1)      // 删除key=1的键值对
	// fmt.Println(m)    // map[]

	// mm := make(map[int]map[int]string)
	// val2_1, ok := mm[2][1]  // val2_1获取嵌套map中key=2的map中key=1的字符串,ok获取的是该值是否初始化的bool值
	// fmt.Println(val2_1, ok) // false
	// if !ok {
	// 	mm[2] = make(map[int]string)
	// 	fmt.Println("初始化完成") // 初始化完成
	// }
	// mm[2][1] = "hello world!"
	// val2_1, ok = mm[2][1]
	// fmt.Println(val2_1, ok) // hello world! true

	// 迭代操作
	// sm := make([]map[int]string, 5)
	// for _, v := range sm {
	// 	v = make(map[int]string, 1)
	// 	v[1] = "OK"
	// 	fmt.Println(v) // 迭代5次,都输出map[1:OK]
	// }
	// fmt.Println(sm) // [map[] map[] map[] map[] map[]]
	// // 观察以上的输出,可见for循环中获取的v只是sm中value的一个拷贝
	// // 如果要改变sm,需要用索引来操作
	// for k, _ := range sm {
	// 	sm[k] = make(map[int]string, 1)
	// 	sm[k][1] = "OK"
	// 	fmt.Println(sm[k]) // map[1:OK]
	// }
	// fmt.Println(sm) // [map[1:OK] map[1:OK] map[1:OK] map[1:OK] map[1:OK]]

	m_sort := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m_sort))
	i := 0
	for k, _ := range m_sort {
		s[i] = k
		i++
	}
	fmt.Println(s) // [4 5 1 2 3] 每次输出都未必相同
	sort.Ints(s)   // 导入sort包进行排序
	fmt.Println(s) // [1 2 3 4 5]
}
