## 接口interface

- 接口是一个或多个方法签名的集合
- 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口，这称为Structural Typing
- 接口只有方法声明，没有实现，没有数据字段
- 接口可以匿名嵌入其他接口，或嵌入到结构中
- 将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
- 只有当接口存储的类型和对象都为nil时，接口才等于nil
- 接口调用不会做receiver的自动转换
- 接口同样支持匿名字段方法
- 接口也可实现类似OOP中的多态
- 空接口可以作为任何类型数据的容器



------

## 接口的声明与使用

- 接口是一个或多个方法签名的集合
- 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口，这称为Structural Typing。<font color="#ff0000">所以任何一个类型都实现了空接口</font>，就像其它语言都继承Object类一样

```go
// 声明接口——一个或多个方法签名的集合
type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

// struct类型拥有所有方法签名，即实现该接口
func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func Disconnect1(usb USB) {
	fmt.Println("我是接口类型")
}

func main() {
	// 定义接口变量
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect() // Connect: PhoneConnecter

	b := PhoneConnecter{"PhoneConnecter"}
	fmt.Printf("b type:%T\n", b) // b type:main.PhoneConnecter
	// 判断是否接口类型
	Disconnect1(b) // Disconnected 说明是USB类型
}
```



------

## 嵌入接口

```go
type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}
// 其余按照上面实现即可
```



------

## 类型断言

- 通过类型断言的ok pattern可以判断接口中的数据类型
- 使用type switch则可以针对空接口进行比较全面的类型判断

```go
// 声明接口——一个或多个方法签名的集合
type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

// struct类型拥有所有方法签名，即实现该接口
func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func Disconnect1(usb USB) {
	// 可以通过判断usb的结构来确定
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected1:", pc.name)
		return
	}
	fmt.Println("Unknown device.")
}

// 形参为空接口，即可以传递任何类型
func Disconnect2(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected2:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

func main() {
	// 定义接口变量
	b := PhoneConnecter{"PhoneConnecter"}
	// 判断实现接口的变量的类型
	Disconnect1(b) // Disconnected1: PhoneConnecter
	Disconnect2(b) // Disconnected2: PhoneConnecter
}
```



------

## 接口转换

- 可以将拥有超集的接口转换为子集的接口，以上代码的USB接口可以转换为Connecter接口，成立的前提是Connecter接口嵌入在USB接口中。

```go
// 接口的向下转换
var c Connecter
c = Connecter(a)
c.Connect() // Connect: PhoneConnecter
```



------

## 只有当接口存储的类型和对象都为nil时，接口才等于nil

```go
var d interface{}
fmt.Println(d == nil) // true

var p *int = nil
d = p
fmt.Println(d == nil) // false
```

