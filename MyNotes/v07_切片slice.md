## 切片Slice

- 其本身并不是数组，它指向底层的数组
- 作为变长数组的替代方案，可以关联底层数组的局部或全部
- 为引用类型
- 可以直接创建或从底层数组获取生成
- 一般使用`make()`创建
- 使用`len()`获取元素个数，cap()获取容量
- 如果多个`slice`指向相同底层数组，其中一个的值改变会影响全部

- `make([]T, len, cap)`
- 其中`cap`可以省略，则和`len`的值相同
- `len`表示存放元素个数，`cap`表示容量



------

## 创建slice

```go
func main() {
	// 1.直接创建slice
	var s1 []int
	fmt.Println(s1) // []

	// 2.从底层数组获取slice
	// 首先创建数组
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a) // [1 2 3 4 5 6 7 8 9 0]
	// 截取数组创建slice
	s2 := a[5:10]   // 注意索引值的含义：5代表从下表为5的元素开始截取，取几个元素呢，由(右值-左值)决定
	fmt.Println(s2) // [6 7 8 9 0]

	s3 := a[5:]     // 取数组的后5个元素
	fmt.Println(s3) // [6 7 8 9 0]

	s4 := a[:5]     // 取数组的前5个元素
	fmt.Println(s4) // [1 2 3 4 5]
    
    s := a[:] // 创建一个完整数组的切片
	fmt.Println(s) // [1 2 3 4 5 6 7 8 9 0]

    // make([]T, len, cap)创建数组
	s5 := make([]int, 3)          // 长度len=3,未定义容量cap
	fmt.Println(len(s5), cap(s5)) // 输出3 3，说明如果未定义容量，则cap=len
	fmt.Println(s5)               // [0 0 0]

	s6 := make([]int, 3, 6) // len=3,cap=10
	fmt.Println(s6)         // [0 0 0]
}
```



----

## slice与底层数组的对应关系

* 如果两个不同的切片截取同一个底层数组，那它们指向同一个数组，对应元素的地址相同

* 长度 = 截取的右边界 - 截取的左边界

* 容量 = 数组的长度 - 截取的左边界

  ```go
  c := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}
  sa := c[3:7]
  fmt.Println("sa:", string(sa), "len:", len(sa), "cap:", cap(sa)) // sa: defg len: 4 cap: 9
  sb := c[4:6]
  fmt.Println("sb:", string(sb), "len:", len(sb), "cap:", cap(sb)) // sb: ef len: 2 cap: 8
  fmt.Println("sa中e的地址：", &sa[1]) // sa中e的地址： 0xc000058084
  fmt.Println("sb中e的地址：", &sb[0]) // sb中e的地址： 0xc000058084
  ```



----

## 切分slice

- 对slice进行切分会生成一个新的slice变量，新slice和原slice指向同一个底层数组，只不过指向的起始位置可能不同，长度及容量可能也不相同。

  - 当从左边界有截断时，会改变新切片容量大小
  - 当然，因为指向同一个底层数组，对新slice的操作会影响到原来的slice

- 其实新slice变量切分原slice的效果与指向数组的效果相同

- 切分时索引以被切分的slice的切片为准

- 索引不可以超过被切分的slice的容量

- 索引越界不会导致底层数组的重新分配，而是引发错误

  ```go
  // 切分slice
  sc := c[4:6]
  fmt.Println("sc:", string(sc), "len:", len(sc), "cap:", cap(sc)) // sc: ef len: 2 cap: 8
  sd := sc[1:5]
  fmt.Println("sd:", string(sd), "len:", len(sd), "cap:", cap(sd)) // sd: fghi len: 4 cap: 7
  ```



-----

## 切片的append操作

- 可以在slice尾部追加元素

- 可以将一个slice追加在另一个slice尾部

- 如果当前slice追加元素后：

  - 最终长度未超过该slice的容量，则返回该slice；
    - 但容量不变，指向的地址不变
  - 最终长度如果超过了，则重新分配数组并拷贝原始数据。
    - 数组长度（即切片容量）每次成倍增加。
    - 指向的地址发生改变，因此对追加后的slice的操作不会影响到原slice

  ```go
  // slice的append操作
  se := make([]int, 3, 6)
  fmt.Printf("%p\n", se) // 0xc0000103c0
  se = append(se, 1, 2, 3)
  fmt.Printf("%v %p\n", se, se) // [0 0 0 1 2 3] 0xc0000103c0
  se = append(se, 4, 5, 6)
  fmt.Printf("%v %p\n", se, se) // [0 0 0 1 2 3 4 5 6] 0xc00000a120
  ```

  

------

## copy操作

- copy函数：`copy(dst,src)`

- 假设存在切片变量`s1,s2`,那么`copy(s1,s2)`的结果会根据`s1,s2`的长度而定

  - `len(s1) > len(s2)`:s2会将自己的元素替换掉s1对应索引的元素，s1其他索引的元素不变

    ```go
    s1 := []int{1, 2, 3, 4, 5, 6}
    s2 := []int{7, 8, 9}
    copy(s1, s2)
    fmt.Println("s1:", s1, "s2:", s2) //s1: [7 8 9 4 5 6] s2: [7 8 9]
    ```

  - `len(s1) < len(s2)`:s2会将自己的元素替换掉s1对应索引的元素，直到s1的最后一个元素

    ```go
    s1 := []int{7, 8, 9}
    s2 := []int{1, 2, 3, 4, 5, 6}
    copy(s1, s2)
    fmt.Println("s1:", s1, "s2:", s2) //s1: [1 2 3] s2: [1 2 3 4 5 6]
    ```

- 总结：`copy(dst,src)`会将`src`的元素替换掉`dst`对应索引的元素，直到长度短的slice的最后一个元素。







