## 常见的并发模式

### 生产者/消费者模式

并发编程中最常见的例子就是生产者/消费者模式，该模式通过平衡生产线程和消费线程的工作能力来提高程序的整体处理数据的速度。

简单地说，就是生产者生产一些数据，然后放到成果队列中；同时消费者从成果队列中来取这些数据。这样就让生产消费变成了异步的两个过程。

```go
// 生产者，输出factor整数倍的序列
func Producer(factor int, out chan<- int) {
  for i := 0; ; i++ {
    out <- i * factor
  }
}

// 消费者
func Consumer(in <-chan int) {
  for v := range in {
    fmt.Println(v)
  }
}
```

### 发布/订阅模式
