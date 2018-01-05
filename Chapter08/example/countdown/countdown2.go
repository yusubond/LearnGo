// Name: countdown2.go
// Func: 按下回车键时，取消火箭发送
// Author: subond
// Date: 5 Jan, 2018
package main

import "fmt"
import "time"
import "os"

func main() {
  // 添加abort功能，按下回车键时取消发射火箭
  abort := make(chan struct{})
  // 创建一个goroutine，负责读取单个字符
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

func launch() {
  fmt.Println("Lift Off!")
}
