## 多路复用select

select语句的通用形式如下所示：

```go
select {
case <-ch1:
  // ...
case x := <-ch2:
  // ...
case ch3 <- y:
  // ...
default:
  // ...
}
```

`select`语句与switch语句很像，它有一系列的情况和一个可选的默认分支。*每一个情况指定一次通信(在一些通道上进行发送和接收操作)和关联的一段代码块*。接收表达式操作可以出现在它本身上，如第一种情况`case <-ch1:`；或者在一个短变量声明中，如第二种情况，这种情况下可以直接引用所接收的值。

select一直等待，直到一次通信来告诉它有一些情况可以执行，然后执行这次通信，执行此情况下对应的代码块；其他通信将不会发生。值得注意的是，对于没有对应情况的select，`select{}`将永远等待。

下面是一个小练习，演示火箭发射的过程，按下回车键取消发射。

```go
// countdown2.go
func main() {
  // 添加abort功能，按下回车键时取消发射火箭
  abort := make(chan struct{})
  // 创建一个goroutine，负责读取单个字符
  // go func() 以并发的方式运行匿名函数
  go func() {
    os.Stdin.Read(make([]byte, 1))
    abort <- struct{}{}
  }()

  fmt.Println("Commencing countdown. Press return to abort")
  select {
  case <-time.After(10 * time.Second):
  case <-abort:
    fmt.Println("Launch aborted!")
    return
  }
  launch()
}
```

`time.After`函数立即返回一个通道，然后启动一个新的goroutine在间隔指定时候后，发送一个值到它上面。在上面的程序中time.After将在10s后发送一个值到第一个case。

综合来看，上面的select语句等待两个事件中第一个到达的事件，中止事件或指示事件过去10s的事件。

```go
// 偶数时发送，奇数时接收
ch := make(chan int, 1)
for i := 0; i < 10; i++ {
  select {
  // 奇数时接收
  case x := <-ch:
    fmt.Println(x)
  // 偶数时发送
  case ch <- i:
  }
}
```

上面的程序中，因为通道ch的缓冲区大小为1，所以要么是空的，要么是满的，只能在一种情况下可以执行。

🌰：[evennumber.go](example/countdown/evennumber.go)

**如果多个情况同时满足，seletc则随机选择一个执行**。

下面看一下select的默认情况的使用，它用来指定在没有其他通信发生时可以立即执行的动作。

```go
select {
case <-abort:
  fmt.Println("Launch aborted!")
  return
default:
  // ...
}
```
这个例子中，select尝试从abort通道中接收一个值，如果没有值，它什么也不做。这是一个非阻塞的接收操作。
