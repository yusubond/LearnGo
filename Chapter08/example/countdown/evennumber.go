// Name: evennumber.go
// Func: 只发送偶数数字
// Author: subond
// Date: 5 Jan, 2018
package main

import "fmt"

func main() {
  ch := make(chan int, 1)
  for i := 0; i < 10; i++ {
    select {
    case x := <-ch:
      fmt.Println(x)
    case ch <-i:
    }
  }
}
