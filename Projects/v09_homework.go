package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i =", i) // 仅获取值

		defer func() {
			fmt.Println("defer_closure i =", i) // 匿名函数获取外部变量的内存地址
		}()

		fs[i] = func() {
			fmt.Println("closure i =", i) // 匿名函数获取外部变量的内存地址
		}
	}

	for _, f := range fs {
		f()
	}

	// 输出如下
	/*
		closure i = 4
		closure i = 4
		closure i = 4
		closure i = 4
		defer_closure i = 4
		defer i = 3
		defer_closure i = 4
		defer i = 2
		defer_closure i = 4
		defer i = 1
		defer_closure i = 4
		defer i = 0
	*/
	/*
		分析：重点把握匿名函数/闭包获取外部变量的内存地址和defer语句的调用顺序两个关键点
		- 其中两句defer语句是在最后才执行的，重点在于分析函数数组fs的匿名函数。
		- 匿名函数调用外部变量是获取其内存地址，每次循环创建的匿名函数都存进fs数组中。
		- 然后for循环结束，i自增为4，执行到for range语句，执行fs数组内的每个函数，都输出i=4
		- defer语句只要把握其特点即可，不需要再分析
	*/

}
