// Name: server.go
// Func: 一个简单的web服务器，同时返回请求的数量
// Author: subond
// Date: Dec 11, 2017
package main

import (
  "fmt"
  "log"
  "net/http"
  "sync"
)

var mu sync.Mutex
var count int

func main() {
  http.HandleFunc("/", handler)  // echo处理程序
  http.HandleFunc("/count", counter) // counter处理函数
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// echo处理程序
// 输入 http.Request
func handler(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  count++
  mu.Unlock()
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// 回显目前的调用次数
func counter(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  fmt.Fprintf(w, "Count %d\n", count)
  mu.Unlock()
}
