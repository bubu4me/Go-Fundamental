## 数组Array

- 定义数组的格式有多种，将在以下通过代码一一演示

- 数组长度也是类型的一部分，因此具有不同长度的数组为不同类型

- 注意区分指向数组的指针和指针数组

- 数组在Go中为值类型

- 数组之间可以使用==或\!=进行比较，但不可以使用<或>

- 可以使用new来创建数组，此方法返回一个指向数组的指针

- Go支持多维数组

  ```go
  func main(){
      // 数组
      a_int := [3]int{} // 3为数组长度，大括号没有值默认为0
      a_str := [3]string{} // 3为数组长度，大括号没有值默认为空字符串
      b := [3]int{1,2,3} // b[0]=1, b[1]=2, b[2]=3
      c := [100]int{19:1} // c[19]=1,其余元素均为0
      d := [...]int{1,2,3,4,5} // ...会根据大括号内值的个数计算长度
      e := [...]int{19:1} // b[19]=1,根据索引最大值计算，长度为20
      
      // 数组指针
      a := [...]int{99: 1}
  	var ap *[100]int = &a // 此处的理解：将[100]int看作一种类型，然后声明此类型的指针，这个指针指向该类型的对象a的地址
      // ap := &a // 简写
  	fmt.Println(p)
      /* 输出：
      &[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
      */
      
      
      // 指针数组
      x, y := 1, 2
      pa := [...]*int{&x, &y} // 此处的理解：将*int看作一种类型，然后声明此类型的数组
      fmt.Println(pa)
      /* 输出：
      [0xc000058080 0xc000058088]
      */
      
      // 数组比较是否相等，只有类型和长度都相同的数组才能比较
      a1 := [...]int{1, 2}
  	a2 := [2]int{1, 2}
  	a3 := [...]int{1, 1}
  	fmt.Println(a1 == a2) // true
  	fmt.Println(a1 == a3) // false
  	fmt.Println(a2 == a3) // false
      
      // new创建指向数组的指针
      p := new([10]int)
  	fmt.Println(p) // &[0 0 0 0 0 0 0 0 0 0]
      
      // 多维数组
      // 注意：多维数组中只有顶级数组的长度可以用...来代替，由系统计算，非顶级数组不可以，即[2][3]中2可以用...代替
      double_a := [2][3]int{
  		{1, 1, 1},
  		{2, 2, 2}} // 注意，结尾的大括号必须这样写，否则报错
  	fmt.Println(double_a) // [[1 1 1] [2 2 2]]
  }
  ```

