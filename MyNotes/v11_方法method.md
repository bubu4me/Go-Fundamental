## 方法method

- Go中虽然没有class，但依旧有method
- 通过显示说明receiver来实现与某个类型的组合
- 只能为同一个包中的类型定义方法
- receiver可以是类型的值或者指针
- 不存在方法重载
- 可以使用值或指针来调用方法，编译器会自动完成转换
- 从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的第1个参数(Method Value vs. Method Expression)
- 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
- 类型别名不会拥有底层类型所附带的方法
- 方法可以调用结构中的非公开字段



------

## 使用值或指针来调用方法

- Go中通过在方法中声明receiver的类型，来实现对同名方法的不同调用
- 不存在方法重载，但可以使用值或指针来调用方法，编译器会自动完成转换

```go
type A struct {
	Name string
}

type B struct {
	Name string
}

// 定义方法：通过显示说明receiver来实现与某个类型的组合
// 此处的receiver：(a A)
func (a A) Print() {
	a.Name = "aa"
	fmt.Println("A", a.Name)
}

// 此处的receiver：(b *B)
func (b *B) Print() {
	b.Name = "bb"
	fmt.Println("B", b.Name)
}

func main() {
	a := A{}
	a.Print()                   // A aa
	fmt.Println("Name", a.Name) // Name

	b := B{}
	b.Print()                   // B bb
	fmt.Println("Name", b.Name) // Name bb
}
```



## 类型别名调用方法

```go
type typeInt int

func (a *typeInt) Print() {
	fmt.Println("typeInt")
}

func main() {
	var a typeInt
    // 这种调用称为Method Value
	a.Print()            // typeInt
    
    // 这种调用称为Method Expression
	(*typeInt).Print(&a) // typeInt
}
```



------

## 结构内非公开字段的访问

- 在Go中，私有字段的访问仅限于本包

```go
type C struct {
	name string
}

func (c *C) Print() {
	c.name = "123"
	fmt.Println(c.name)
}

func main(){
    c := C{}
	c.Print()           // 123
	fmt.Println(c.name) // 123
}	
```



------

## 课堂作业

- 根据为结构增加方法的知识，尝试声明一个底层类型为int的类型，并实现调用某个方法就递增100

```go
type A struct {
	int
}

func (a *A) Increase(num int) {
	a.int += num
}

// 优化
type TZ int

func (tz *TZ) Increase(num int) {
	*tz += TZ(num) // 虽然*tz和num底层都是int类型，但是相加时仍然要通过类型转换为相同类型
}

func main() {
	a := A{}
	a.Increase(100)
	fmt.Println(a) // {100}

	var t TZ
	t.Increase(100)
	fmt.Println(t) // 100
}

```

