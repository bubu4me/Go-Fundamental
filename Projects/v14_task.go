package main

import (
	"fmt"
	// "time"
)

func main() {
	// go Go()                     // 单独运行此语句没有任何输出
	// time.Sleep(2 * time.Second) // Go Go Go!!!

	c := make(chan bool) // 创建一个可以发送bool类型的channel
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true // 存入c这个channel，发送数据为true
		close(c)
	}()
	// <-c // 执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。

	// 对channel进行迭代操作
	for v := range c {
		fmt.Println(v)
	}
	// 输出：
	/*
		fatal error: true
		all goroutines are asleep - deadlock!
	*/
	// 注意：在对channel进行迭代操作之前，必须先明确对该channel关闭成功，否则程序变成死锁，最后崩溃退出
	// 在line 15之后关闭channel，line 22输出true
}

func Go() {
	fmt.Println("Go Go Go!!!")
}
