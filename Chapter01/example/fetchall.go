// Name: fetchall.go
// Func: ；并发获取多个url
// Author: subond
// Date: Dec 11, 2017
package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
  "strings"
)

func main() {
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    // 启动一个goroutine
    go fetch(url, ch)
  }
  for range os.Args[1:] {
    // 从通道ch接收数据
    fmt.Println(<-ch)
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// function fetch
func fetch(url string, ch chan<- string) {
  start := time.Now()
  if !strings.HasPrefix(url, "http://") {
    url = "http://" + url
  }
  resp, err := http.Get(url)
  if err != nil {
    // 如果出错，发送错误信息至通道ch
    ch <- fmt.Sprint(err)
    return
  }
  // 拷贝到ioutil.Discard输出流进行丢弃
  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close()
  if err != nil {
    ch <- fmt.Sprint("while reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
