// Name: strchan.go
// Func: 缓冲通道
// Author: subond
// Date: Dec 26, 2017
package main

import "fmt"

func main() {
  // 创建一个chan string类型，容量为3的缓冲通道
  ch := make(chan string, 3)
  ch <- "A"
  ch <- "B"
  ch <- "C"
  fmt.Printf("缓冲通道的初始状态\n")
  fmt.Printf("当前通道的容量为：%d\n", cap(ch))
  fmt.Printf("当前通道的长度为：%d\n", len(ch))
  fmt.Println()
  fmt.Printf("接收一个元素为：%s\n", <-ch)
  fmt.Printf("当前通道的容量为：%d\n", cap(ch))
  fmt.Printf("当前通道的长度为：%d\n", len(ch))
  fmt.Println()
  fmt.Printf("接收一个元素为：%s\n", <-ch)
  fmt.Printf("当前通道的容量为：%d\n", cap(ch))
  fmt.Printf("当前通道的长度为：%d\n", len(ch))
}
