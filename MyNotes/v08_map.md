## map

- 类似其他语言中的哈希表或者字典，以key-value形式存储数据
- Key必须是支持==或\!=比较运算的类型，不可以是函数、map或slice
- Map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
- Map使用make()创建，支持:=简写方式
- 可以嵌套使用

- make([keyType]valueType, cap)，cap为容量，可忽略
- 超出容量时会自动扩容，但尽量提供一个合理的初始值
- 使用len()获取元素个数

- 键值对不存在时自动添加，使用delete()删除某键值对
- 使用 for range 对map和slice进行迭代操作



------

## map的创建方法

```go
// map的创建方式
// 第一种
var m1 map[int]string
m1 = map[int]string{} // 可与上一句合并
fmt.Println(m1) // map[]
// 以上的简写形式
m3 := map[int]string{}
fmt.Println(m3) // map[]

// 第二种
var m2 map[int]string
m2 = make(map[int]string) // 可与上一句合并
fmt.Println(m2) // map[]
// 以上的简写形式
m4 := make(map[int]string)
fmt.Println(m4) // map[]
```



## map的简单使用与嵌套使用

```go
// 展示添加、删除键值对，以及获取key的对应value的操作
m := make(map[int]string)
m[1] = "hello"    // 添加key-value
fmt.Println(m)    // map[1:hello]
val1 := m[1]      // 取出key=1的value
fmt.Println(val1) // hello
delete(m, 1)      // 删除key=1的键值对
fmt.Println(m)    // map[]

// map的嵌套使用，以此类推
mm := make(map[int]map[int]string)
val2_1, ok := mm[2][1]  // val2_1获取嵌套map中key=2的map中key=1的字符串,ok获取的是该值是否初始化的bool值：是则true，否则false
fmt.Println(val2_1, ok) // false
if !ok {
    mm[2] = make(map[int]string) // 进行初始化
    fmt.Println("初始化完成") // 初始化完成
}
mm[2][1] = "hello world!"
val2_1, ok = mm[2][1]
fmt.Println(val2_1, ok) // hello world! true
```



## map的迭代操作

```go
// slice的迭代
for i,v := range slice{ // i是返回的索引，v是返回的值
    
}

// map迭代操作
sm := make([]map[int]string, 5)
for _, v := range sm {
    v = make(map[int]string, 1)
    v[1] = "OK"
    fmt.Println(v) // 迭代5次,都输出map[1:OK]
}
fmt.Println(sm) // [map[] map[] map[] map[] map[]]
// 观察以上的输出,可见for循环中获取的v只是sm中value的一个拷贝
// 如果要改变sm,需要用索引来操作
for k, _ := range sm {
    sm[k] = make(map[int]string, 1)
    sm[k][1] = "OK"
    fmt.Println(sm[k]) // map[1:OK]
}
fmt.Println(sm) // [map[1:OK] map[1:OK] map[1:OK] map[1:OK] map[1:OK]]
```



## map的间接排序

```go
m_sort := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
s := make([]int, len(m_sort))
i := 0
for k, _ := range m_sort {
    s[i] = k
    i++
}
fmt.Println(s) // [4 5 1 2 3] 每次输出都未必相同
sort.Ints(s)   // 导入sort包进行排序
fmt.Println(s) // [1 2 3 4 5]
```



------

## 课堂作业

- 根据在for range部分讲解的知识，尝试将类型为map[int]string的键和值进行交换，变为类型为map[string]int

  ```go
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
  	m2 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
  	fmt.Println(m1) // map[1:a 2:b 3:c 4:d 5:e]
  	fmt.Println(m2) // map[a:1 b:2 c:3 d:4 e:5]
  
  	m3 := make(map[string]int, 5)
  	for k, v := range m1 {
  		m3[v] = k
  	}
  	fmt.Println(m3) // map[a:1 b:2 c:3 d:4 e:5]
  }
  ```

  













