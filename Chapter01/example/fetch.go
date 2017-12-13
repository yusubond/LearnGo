// Name: fetch.go
// Fucn: 模仿curl指令
// Author: subond
// Date: Dec 11, 2017
package main

import (
  "fmt"
  "io/ioutil"
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
    data, err := ioutil.ReadAll(resp.Body)
    // 关闭Body数据流以避免资源泄漏
    resp.Body.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
      os.Exit(1)
    }
    fmt.Printf("%s", data)
  }
}
