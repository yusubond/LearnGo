// Name: crawl2.go
// Func: 并发web爬虫，解决无限制并发以及程序不终止问题
// Author: subond
// Date: 5 Jan, 2018

package main

import (
  "fmt"
  "log"
  "os"
  "links"
)

// crawl2
// 令牌是一个计数信号量，确保并发请求限制在20个以内
var tokens = make(chan struct{}, 20)
func crawl(url string) []string {
  fmt.Println(url)
  // 获取令牌
  tokens <- struct{}{}
  list, err := links.Extract(url)
  // 释放令牌
  <-tokens
  if err != nil {
    log.Print(err)
  }
  return list
}

func main() {
  worklist := make(chan []string)
  var n int  // 等待发送到任务列表的数量，从命令行参数开始算起
  n++
  go func() { worklist <- os.Args[1:] }()

  seen := make(map[string]bool)
  for ; n > 0; n-- {
    list := <-worklist
    for _, link := range list {
      if !seen[link] {
        seen[link] = true
        n++
        go func(link string) {
          worklist <- crawl(link)
        }(link)
      }
    }
  }
}
