// Name: pipeline3.go
// Func: 单向通道类型
// Author: subond
// Date: Dec 26, 2017
package main

import "fmt"

// counter只接收输出通道参数
func counter(out chan<- int) {
  for x := 0; x < 100; x++ {
    out <- x
  }
  close(out)
}
// square接受两个参数，输入通道和输出通道
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
func main() {
  naturals := make(chan int)
  squarers := make(chan int)
  go counter(naturals)
  go squarer(squarers, naturals)
  printer(squarers)
}
