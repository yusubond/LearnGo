// Name: fetch2.go
// Exam: 练习1.7
// Fucn: 使用io.Copy()复制响应内容
// Author: subond
// Date: Dec 11, 2017
package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
)

func main() {
  for _, url := range os.Args[1:] {
    resp, err := http.Get(url)
    if  err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }
    io.Copy(os.Stdout, resp.Body)
    // 关闭Body数据流以避免资源泄漏
    resp.Body.Close()
  }
}
