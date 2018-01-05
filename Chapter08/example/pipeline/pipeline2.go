// Name: pipeline2.go
// Func: 管道，通道连接goroutine，一个goroutine的输出是另一个goroutine的输入，称为管道(pipeline)
// Author: subond
// Date: Dec 26, 2017
package main

import "fmt"

func main() {
  naturals := make(chan int)
  squares := make(chan int)

  // counter
  go func() {
    for x := 0; x < 100; x++ {
      naturals <- x
    }
    // 关闭naturals通道
    close(naturals)
  }()

  // square
  go func() {
    for x := range naturals {
      squares <- x * x
    }
    // 关闭squares通道
    close(squares)
  }()

  // printer
  for x := range squares {
    fmt.Println(x)
  }
}
