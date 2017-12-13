// Name: fetch3.go
// Exam: 练习1.8
// Func: 使用strings.HasPrefix函数检查url中是否包含"http://"前缀
// Author: subond
// Date: Dec 11, 2017
package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "strings"
)

func main() {
  for _, url := range os.Args[1:] {
    if strings.HasPrefix(url, "http://") == false {
      url = "http://" + url
    }
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
