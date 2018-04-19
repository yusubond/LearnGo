//  Name: producerandconsumer.go
// Func: 并发编程，生产者与消费者模式
// Author: subond
// Date: Apr 28, 2018
package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
)
// 生产者
func Producer(factor int, out chan<- int) {
  for i := 0; i < 10 ; i++ {
    out <- i * factor
  }
}
// 消费者
func Consumer(in <-chan int) {
  for v := range in {
    fmt.Println(v)
  }
}

func main() {
  ch := make(chan int, 4)
  go Producer(3, ch)
  go Producer(5, ch)
  go Consumer(ch)
  //Ctrl+C退出
  sig := make(chan os.Signal, 1)
  signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
  fmt.Printf("Quit (%v)\n", <-sig)
}
