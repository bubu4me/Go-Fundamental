# Go编程基础

## Go程序的一般结构

* Go程序是通过package来组织的（与python类似）
* 只有package名称为main的包可以包含main函数
* 一个可执行程序有且只有一个main包

* 导入非main包：import
* 定义常量：const
* 全局变量的声明与赋值：var
* 结构或接口的声明：type
* 函数声明：func

```go
// 当前程序的包名
package main

// 导入其他包
import "fmt"

// 常量的定义
const PI = 3.14

// 全局变量的声明与赋值
var name = "gopher"

// 一般类型的声明
type newType int

// 结构的声明
type gopher struct{
	// 一些变量
}

// 接口的声明
type golang interface{}

// 由main函数作为程序入口点启动
func main {
	fmt.Println("hello world!")
}

```

------

* package 别名

  * 当使用第三方包时，包名可能会非常接近或者相同，此时可以使用别名来进行区别和调用

  ```go
  import (
  	io "fmt"//使用别名为io
  )
  
  func main{
      // 使用别名调用包
      io.Println("hello world")
  }
  ```

------

* 可见性规则
  * Go语言中，使用<font color="#ff00ff">大小写</font>来决定该 常量、变量、类型、接口、结构或函数 是否可以被外部包调用：
    * 根据约定，<font color="#00AAff">函数名首字母小写为private，大写为public</font>

------

* 多个常量、全局变量或一般类型（非接口、非结构）的声明方法

  ```go
  const (
  	PI = 3.14
  	const1 = "1"
  	const2 = 2
  )
  
  var (
  	name = "gopher"
      name1 = "1"
      name2 = 2
  )
  
  type(
  	newType int
      type1 float32
      type2 string
      type3 byte
  )
  ```

  

