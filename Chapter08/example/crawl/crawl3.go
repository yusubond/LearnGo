// Name: crawl3.go
// Func: 并发web爬虫，使用固定的20个长期存活的goroutine来避免过度并发
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
  worklist := make(chan []string)   // 可能有重复的URL任务列表
  unseenlinks := make(chan string)  // 去重复的URL任务列表

  // 创建一个goroutine，负责从命令行参数获取初始任务列表
  go func() { worklist <- os.Args[1:] }()

  // 创建20个爬虫goroutine来获取每个不可见链接
  for i := 0; i < 20; i++ {
    go func() {
      for link := range unseenlinks {
        foundlinks := crawl(link)
        go func() { worklist <- foundlinks }()
      }
    }()
  }
  // 主goroutine对url进行去重，并把没有爬取过的条目发送给爬虫程序
  seen := make(map[string]bool)
  for list := range worklist {
    for _, link := range list {
      if !seen[link] {
        seen[link] = true
        unseenlinks <- link
      }
    }
  }
}
