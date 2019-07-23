package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// https://gowalker.org/time#_constants
	// 不要随便修改常量值
	fmt.Println(t.Format(time.ANSIC))
}
