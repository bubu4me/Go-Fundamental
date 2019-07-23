package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	// for _, v := range s {
	// 	go func() {
	// 		fmt.Println(v)
	// 	}()
	// }
	// select {}
	// 输出：
	/*
		c
		c
		c
		fatal error: all goroutines are asleep - deadlock!
	*/
	/*
		分析：
		死锁是故意为之，不必理会
		输出3个都是c，可见闭包中引用了v的地址，那么如何进行值传递呢？
		作为参数传入v即可,但发现输出的顺序与字符串存储的索引不一致
	*/
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	select {}
}
