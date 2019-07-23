package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello World!")
}

func main() {
	// u := User{1, "OK", 12}
	// Info(u)
	/*
		Type: User
		Fields:
		    Id: int = 1
		  Name: string = OK
		   Age: int = 12
		 Hello: func(main.User)
	*/
	// Info(&u)
	/*
		Type: User
		传入的参数非结构类型
	*/

	// 匿名字段信息
	// m := Manager{User: User{1, "OK", 12}, title: "123"}
	// t := reflect.TypeOf(m)
	// fmt.Printf("%#v\n", t.Field(0))                  // reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x4ef520), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
	// fmt.Printf("%#v\n", t.Field(1))                  // reflect.StructField{Name:"title", PkgPath:"main", Type:(*reflect.rtype)(0x4d7f00), Tag:"", Offset:0x20, Index:[]int{1}, Anonymous:false}
	// fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0})) // reflect.StructField{Name:"Id", PkgPath:"", Type:(*reflect.rtype)(0x4d7780), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	// fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1})) // reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4d7f00), Tag:"", Offset:0x8, Index:[]int{1}, Anonymous:false}

	// 利用反射修改对象状态
	// x := 123
	// v := reflect.ValueOf(&x)
	// v.Elem().SetInt(999)
	// fmt.Println(x) // 999

	// 自定义修改方法
	// u1 := User{1, "OK", 2}
	// Set(&u1)
	// fmt.Println(u1) //{1 BYEBYE 2}

	// 通过反射"动态"调用方法
	u2 := User{1, "OK", 12}
	v := reflect.ValueOf(u2)
	mv := v.MethodByName("Hello2")

	args := []reflect.Value{reflect.ValueOf("bubu"), reflect.ValueOf(11)} // 设置修改值，要按照所获得方法的形参顺序逐一设置
	mv.Call(args)                                                         // Hello bubu , my name is OK , I am 11 years old
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	// 如果传入的是地址，会发生报错
	// 所以先判断传入参数的类型
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("传入的参数非结构类型")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ { // NumField()返回所有字段的数量
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ { // NumField()返回所有方法的数量
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Not settable")
		return
	} else {
		v = v.Elem()   // 返回接口包含的或者指针指向的值
		fmt.Println(v) // {1 OK 2}
	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
	// 如何判断是否包含所要查询的字段呢？
	f1 := v.FieldByName("Name1")
	if !f1.IsValid() {
		fmt.Println("Invalid:", f1) // Invalid: <invalid reflect.Value>
	}
}

func (u User) Hello2(name string, age int) {
	fmt.Println("Hello", name, ", my name is", u.Name, ", I am", age, "years old")
}
