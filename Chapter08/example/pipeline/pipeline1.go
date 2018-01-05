// Name: pipeline1.go
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
    for x := 0; ; x++ {
      naturals <- x
    }
  }()

  // square
  go func() {
    for {
      x := <- naturals
      squares <- x * x
    }
  }()

  // printer
  for {
    fmt.Println(<-squares)
  }
}
