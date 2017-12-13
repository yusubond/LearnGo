// Name: fetch3.go
// Exam: 练习1.9
// Func: 输出HTTP的响应码
// Author: subond
// Date: Dec 11, 2017
package main

import (
  "fmt"
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
    fmt.Println(resp.Status)
  }
}
