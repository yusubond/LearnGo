## channel

目录：

  * [基本概念](#基本概念)
  * [无缓冲通道](#无缓冲通道)
  * [缓冲通道](#缓冲通道)

### 基本概念

如果说goroutine是Go程序并发的执行体，那么通道就是它们之间的连接。**通道是可以让一个goroutine发送特定值到另一个goroutine的通信机制**。

每一个通道是一个具体类型的导管，就做通道的 *元素类型*。一个有int类型元素的通道，写作`chan int`，使用内置函数`make`可以创建一个通道。

```go
ch := make(chan int)      // ch的类型为"chan int"
```

**通道是一个使用`make`函数创建的数据结构的引用**，这一点与map类似。*当复制或作为参数传递到一个函数时，复制的是引用，因此调用者和被调用者都引用同一份数据结构*。通道的零值是nil。

同种类型的通道可以使用`==`符号进行比较，当两者都是同一个通道数据的引用时，返回结果为`true`。

通道有两个主要操作：发送(send)和接收(receive)，统称为通信。send语句从一个goroutine传输一个值到另一个执行接收表达式的goroutine。两个操作都是用`<-`操作符。

发送语句中，通道和值分别放在`<-`的左右两边；接收语句中，`<-`放在通道操作数前面。

```go
ch <- x         // 发送x到通道ch
rx = <- ch      // 从通道ch中接收一个值，并赋值给rx
<- ch           // 接收语句，丢弃结果
```

通道还有第三个操作：关闭(close)，它设置一个标志位来指示当前已经发送完毕，后面没有值了；关闭后的发送操作将导致宕机。其次，在一个已经关闭的通道上进行接收操作，将获取所有已经发送的值，直到通道为空；这时任何接收操作会立即完成，同时获取一个通道元素类型对应的零值。

内置函数`close`可以关闭通道。

```go
close(ch)
```

### 无缓冲通道

无缓冲通道上的发送操作会阻塞，直到另一个goroutine在对应的通道上执行接收操作，这时值传送完成，两个goroutine都可以继续执行。因此，使用无缓冲通道进行通信，将导致发送和接收goroutine同步化，无缓冲通道也称为 **同步信道**。

通道可以用来连接goroutine，这样一个的输出是另一个的输入，这个叫做管道(pipeline)。

栗子：[pipeline1.go](example/pipeline/pipeline1.go)  
栗子：[pipeline2.go](example/pipeline/pipeline2.go)  

注意：结束的时候，关闭每一个通道不是必须的。**只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道**。例如

```go
go func() {
  for x := 0; x < 100; x++ {
    naturals <- x
  }
  // 数据发送完毕，关闭通道
  close(naturals)
}
```

Go语言中还提供了 *单向通道类型*，仅仅导出发送或接受操作。例如，类型`chan<- int`是一个只能发送的通道，允许发送但不允许接受；类型`<-chan int`是一个只能接收的int类型通道，允许接收但是不能发送。`<-`相对于关键字`chan`的位置只是一个帮助记忆的点，但是违反这个原则会在编译时被检查出来。

[pipeline2.go](exmaple/pipeline2.go)的三个函数可以进行如下划分：

```go
func counter(out chan int)
func square(out, in chan int)
func printer(in chan int)
```

下面用单向通道类型，实现上面的三个函数：

```go
// counter只接收输出通道
func counter(out chan<- int) {
  for x := 0; x < 100; x++ {
    out <- x
  }
  close(out)
}
// squarer接收两个参数：输入通道和输出通道
func squarer(out chan<- int, in <-chan int) {
  for v := range in {
    out <- v * v
  }
  close(out)
}
// printer只接收输入通道
func printer(in <-chan int) {
  for v := range in {
    fmt.Println(v)
  }
}
```

### 缓冲通道

缓冲通道有一个元素队列，队列的最大长度在创建的时候通过`make`函数指定，例如，创建一个可以容纳三个字符串的缓冲通道。

```go
ch = make(chan string, 3)
```

缓冲通道上的发送操作在队列的尾部插入一个元素，接收操作从队列的头部移除一个元素。

🌰：[strchan.go](example/strchan.go)
