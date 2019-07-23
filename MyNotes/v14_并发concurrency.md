## 并发concurrency

- 从源码的解析来看，goroutine只是由官方实现的超级“线程池”而已。每个实例4-5KB的栈内存占用和由于实现机制而大幅较少的创建和销毁开销，是制造Go号称的高并发的根本原因。另外，goroutine的简单易用，也在语言层面上给予了开发者巨大的便利。

### 并发不是并行

- 并发主要由切换时间片来实现“同时”运行，而并行则是直接利用多核实现多线程的运行，但Go可以设置使用核数，以发挥多核计算机的能力

### goroutine奉行通过通信来共享内存，而不是共享内存来通信



-----

## Channel

- channel是goroutine沟通的桥梁，大都是阻塞同步的
- 通过make创建，close关闭
- channel是引用类型
- 可以使用for range来迭代不断操作channel
- 可以设置单向或双向通道
- 可以设置缓存大小，在未被填满前不会发生阻塞



------

## select

- 可以处理一个或多个channel的发送与接收
- 同时有多个可用的channel时按随机顺序处理
- 可用空的select来阻塞main函数
- 可设置超时



------

## channel的读取操作

参考博客：http://c.biancheng.net/view/97.html

```go
func main() {
    // 创建一个goroutine
	// go Go()                     // 单独运行此语句没有任何输出
	// time.Sleep(2 * time.Second) // Go Go Go!!!

	c := make(chan bool) // 创建一个可以发送bool类型的channel
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true // 存入c这个channel，发送数据为true
	}()
	<-c // 执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。
}

func Go() {
	fmt.Println("Go Go Go!!!")
}
```



# <font color="red">重看并发的视频</font>











