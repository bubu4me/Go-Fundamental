## 结构struct

* Go中的struct与C中的struct非常相似，并且Go没有class
* 使用type<Name> struct{}定义结构，名称遵循可见性规则
* 支持指向自身的指针类型成员
* 支持匿名结构，可用作成员或定义成员变量
* 匿名结构也可以用于map的值
* 可以使用字面值对结构进行初始化
* 允许直接通过指针来读写结构成员
* 相同类型的成员可进行直接拷贝赋值
* 支持==与\!=比较运算符，但不支持>或<
* 支持匿名字段，本质上是定义了以某个类型名为名称的字段
* 嵌入结构作为匿名字段看起来像继承，但不是继承
* 可以使用匿名字段指针



------

## struct的简单声明和初始化

```go
type person1 struct {
}

type person2 struct {
	Name string
	Age  int
}

func main() {
	p1 := person1{}
	fmt.Println(p1) // {}

	p2 := person2{}
	fmt.Println(p2) // { 0} 前者为空字符串
    
	// 给结构体赋值
	// 方式1
	p2.Name = "Go"
	p2.Age = 9
	fmt.Println(p2) // {Go 9}
	// 方式2
	p3 := person2{
		Name: "Go",
		Age:  9,
	}
	fmt.Println(p3) // {Go 9}
}
```



------

## struct作为参数传递

- strcut作为参数传递是值传递，如果要传递它的内存地址，需要用到指针

```go
func main() {
	p4 := person2{
		Name: "Go",
		Age:  9,
	}
    // 值传递
	fmt.Println(p4) // {Go 9}
	A(p4)           // {Go 11}
	fmt.Println(p4) // {Go 9}
    
	// 内存地址传递
	// 方式1——对变量取地址
	B(&p4)          // &{Go 11}
	fmt.Println(p4) // {Go 11}

	// 方式2——初始化时取地址（推荐）
	p5 := &person2{
		Name: "Go",
		Age:  13,
	}
	fmt.Println(p5) // &{Go 13}
	B(p5)           // &{Go 11}
	fmt.Println(p5) // &{Go 11}
	p5.Name = "go"  // 可直接用取地址的变量名修改结构体的属性
	fmt.Println(p5) // &{go 11}
    
}

func A(per person2) {
	per.Age = 11
	fmt.Println(per)
}

func B(per *person2) {
	per.Age = 11
	fmt.Println(per)
}
```



------

## struct的匿名声明

```go
func main() {
	// 通过匿名结构声明
	p6 := &struct {
		Name string
		Age  int
	}{
		Name: "Go",
		Age:  9,
	}
	fmt.Println(p6) // &{Go 9}
}
```



------

## struct的嵌套与赋值

```go
type person3 struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func main() {
    // 嵌套struct的赋值
	p7 := person3{Name: "Go", Age: 15}
	p7.Contact.Phone = "111111111"
	p7.Contact.City = "beijing"
	fmt.Println(p7)
}
```



------

## struct的匿名字段和赋值

```go
type person4 struct {
	string
	int
}
func main() {
    // struct的匿名字段
    // 赋值顺序要按照匿名字段的顺序，否则报错
    p8 := person4{"Go", 15}
    fmt.Println(p8) // {Go 15}
}
```



------

## struct的嵌入结构

- 重点关注嵌入的结构属性的赋值方法

```go
type human struct {
	Sex string
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
    // 结构体的嵌入
    // 第一种赋值方法
	t := teacher{
		Name:  "t",
		Age:   20,
		human: human{Sex: 1}}
	s := student{
		Name:  "s",
		Age:   10,
		human: human{Sex: 0}}
    fmt.Println(t, s) // {{女} t 20} {{男} s 10}
    // 第二种的赋值方法
	t.human.Sex = "female"
	fmt.Println(t) // {{female} t 20}
	// 第三种的赋值方法
	s.Sex = "male"
	fmt.Println(s) // {{male} s 10}
}
```



------

## 课堂作业

- 如果匿名字段和外层结构有同名字段，应该如何操作？

```go
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
```

