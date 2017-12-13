// Name: server.go
// Func: 一个简单的web服务器
// Author: subond
// Date: Dec 11, 2017
package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", handler)  //echo处理程序
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// echo处理程序
// 输入 http.Request
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
