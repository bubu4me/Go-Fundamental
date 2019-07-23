## 函数function

- Go函数不支持嵌套、重载和默认参数
- 但支持以下特性：
  - 无需声明原型
  - 不定长度变参
  - 多返回值
  - 命名返回值参数
  - 匿名函数
  - 闭包
- 定义函数使用关键字func，且左大括号不能另起一行
- 函数也可以作为一种类型使用





------

## 如何声明函数

```go
func main() {
	A(10)                  // 10
	fmt.Println(B(10))     // 10 10
	fmt.Println(C(10, 11)) // 10 11
	D(1, 2)                // [1 2]
	D(1, 2, 3, 4, 5)       // [1 2 3 4 5]
}

// 定义函数的形式
// 1.带参无返回值函数
func A(a int) {
	fmt.Println(a)
}

// 2.带参有返回值函数，声明返回值名称——推荐
func B(a int) (b, c int) {
	b, c = a, a
	return b, c
}

// 3.带参有返回值函数，只声明返回值类型
func C(a, b int) (int, int) {
	return a, b
}

// 可变长度参数,规定必须放在参数列表中的最右边位置
func D(a ...int) {
	fmt.Println(a) // 可变长度参数相当于一个可变长的数组
}
```



------

## 值传递

* Go中没有引用传递

```go
func main() {
	// 值传递
	a, b := 1, 2
	E(1, 2)                   //函数内： [3 4]
	fmt.Println("函数外：", a, b) // 函数外： 1 2

	// slice的内存地址拷贝
	s1 := []int{1, 2, 3, 4}
	F(s1)                   // 函数内： [5 6 7 8]
	fmt.Println("函数外：", s1) // 函数外： [5 6 7 8]

	// 数据类型的内存地址拷贝
	a1 := 1
	G(&a1)
	fmt.Println("函数外", a1) // 函数外 2
}

// 值传递
func E(a ...int) {
	a[0] = 3
	a[1] = 4
	fmt.Println("函数内：", a) // 函数内： [3 4]
}

// 内存地址拷贝
func F(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Println("函数内：", s) // 函数内： [5 6 7 8]
}

// 数据类型的内存地址拷贝
func G(a *int) {
	*a = 2
	fmt.Println("函数内", *a) // 函数内 2
}
```



------

## 匿名函数和闭包

```go
func main() {
    // 匿名函数
	f1 := func() {
		fmt.Println("匿名函数f1")
	}
	f1() // 匿名函数f1

	// 闭包引用外部变量
	n := 1
	f2 := func() int {
		n++
		return n
	}
	fmt.Println("闭包引用外部变量：", f2()) // 闭包引用外部变量： 2
    fmt.Println(n)                 // 2：n发生改变，可见引用外部变量的实质是传递变量的内存地址

	// 闭包作为函数返回值
	f3 := closure(10)
	fmt.Println(f3(1)) // 11
	fmt.Println(f3(2)) // 12
}

// 闭包
// 观察两次输出的地址，可见闭包中是引用外部变量是传递变量的内存地址
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x) // 0xc0000580a8
	return func(y int) int {
		fmt.Printf("%p\n", &x) // 0xc0000580a8
		return x + y
	}
}
```



------

## defer

- 执行方式类似其它语言中的析构函数，在函数体执行结束后按照调用顺序的<font color="#00ff00">相反顺序</font>逐个执行
- 即使函数发生<font color="#ff0000">严重错误</font>也会执行
- 支持匿名函数的调用
- 常用于资源清理、文件关闭、解锁以及记录时间等操作
- 通过与匿名函数配合可在return之后修改函数计算结果
- 如果函数体内某个变量作为defer匿名函数的参数，则在定义defer时即已经获得了拷贝，否则会引用某个变量的地址

- Go没有异常机制，但有panic/recover模式来处理错误
- panic可以在任何地方引发，但recover只有在defer调用的函数中有效



------

## defer的调用顺序

```go
// defer的调用顺序
fmt.Print("a ")
defer fmt.Print("b ")
defer fmt.Print("c ")
// a c b

for i := 0; i < 3; i++ {
    defer fmt.Print(i, " ") // 2 1 0
}
// 分析：可见defer调用顺序相反
```



-----

## defer匿名函数（闭包）

- 如果函数体内某个变量作为defer匿名函数的参数，则在定义defer时即已经获得了拷贝，否则会引用某个变量的地址

```go
// defer匿名函数（闭包）
for i := 0; i < 3; i++ {
    defer func() {
        fmt.Print(i, " ")
    }() // 加()相当于调用匿名函数
}
// 输出：3 3 3 
// 分析：在闭包调用外部变量实质是获取其内存地址，由于defer调用顺序是相反的，当执行完for循环后i已经自增为3，然后才执行defer的语句，所以输出的都是3
```



------

## panic/recover

- panic一旦引发，就不再执行它之后的语句

```go
func main() {
    // panic
	pA() // Func pA
	pB() // panic: Panic in B
	pC()
}

func pA() {
	fmt.Println("Func pA")
}

func pB() {
	panic("Panic in B")
}

func pC() {
	fmt.Println("Func C")
}
```

- 如何从panic中恢复呢

```go
func main() {
	// panic/recover
	pD() // Func pD
	pE() // Recover in pE
	pF() // Func pF
}

func pD() {
	fmt.Println("Func pD")
}

func pE() {
	defer func() { // 必须放在panic之前，否则panic出现则不再执行该defer语句
		if err := recover(); err != nil {
			fmt.Println("Recover in pE")
		}
	}()
	panic("Panic in pE")
}

func pF() {
	fmt.Println("Func pF")
}
```



------

## 课堂作业

- 观察以下输出结果，分析原因

```go
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
}
```

- 分析：重点把握匿名函数/闭包获取外部变量的内存地址和defer语句的调用顺序两个关键点
  - 其中两句defer语句是在最后才执行的，重点在于分析函数数组fs的匿名函数。
  - 匿名函数调用外部变量是获取其内存地址，每次循环创建的匿名函数都存进fs数组中。
  - 然后for循环结束，i自增为4，执行到for range语句，执行fs数组内的每个函数，都输出i=4
  - defer语句只要把握其特点即可，不需要再分析

