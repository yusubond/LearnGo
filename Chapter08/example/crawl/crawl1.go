// Name: crawl1.go
// Func: 并发web爬虫
// Author: subond
// Date: 5 Jan, 2018

package main

import (
  "fmt"
  "log"
  "os"
  "links"
)

func crawl(url string) []string {
  fmt.Println(url)
  list, err := links.Extract(url)
  if err != nil {
    log.Print(err)
  }
  return list
}

func main() {
  // 创建[]string类型的通道，作为任务列表
  worklist := make(chan []string)

  // 创建一个goroutine，负责从主goroutine接收任务
  // 即，从命令行参数开始
  go func() {worklist <- os.Args[1:] }()

  // 并发爬取web
  seen := make(map[string]bool)
  for list := range worklist {
    for _, link := range list {
      if !seen[link] {
        seen[link] = true
        go func(link string) {
          worklist <- crawl(link)
        }(link)
      }
    }
  }
}
