## 并发的Web爬虫

这是一个简单的网页爬虫，以广度优先的顺序来探索网页的链接图。

目录:

  * [crawl第一版](#crawlv0.1)
  * [crawl第二版](#crawlv0.2)
  * [crawl第三版](#crawlv0.3)

### crawlv0.1

crawl函数

```go
// crawl1
func crawl(url string) []string {
  fmt.Println(url)
  list, err := links.Extract(url)
  if err != nil {
    log.Print(err)
  }
  return list
}
```

main函数

main函数中用一个任务列表记录需要处理的条目队列，每一个条目是一个待爬取的URL列表，并且使用chan来表示队列。每一次对crawl()的调用都发生在它自己的goroutine中，然后将新发现的链接返回任务列表。

```go
func main() {
  // 创建[]string类型的通道，作为任务列表
  worklist := make(chan []string)
  // 创建一个goroutine，负责从主goroutine接收任务
  // 即，从命令行参数开始
  // 这样做的好处就是可以避免死锁
  go func() {worklist <- os.Args[1:] }()
  // 并发爬取web
  seen := make(map[string]bool)
  for list := range worklist {
    for _, link := range list {
      if !seen[link] {
        seen[link] = true
        go func(link string) {
          // craw()函数新发现的链接发送回任务列表
          worklist <- crawl(link)
        }(link)
      }
    }
  }
}
```

注意：发送给任务列表的命令行参数必须在它自己的goroutine中运行以避免死锁。

第一个问题，上面的程序的并行度太高了，以至于到后面的时候超出了程序能打开文件数的限制。

但是，无限制的并行并不是一个好主意，因为系统中总有限制的因素（例如，CPU的核数、磁盘I/O操作磁头和磁盘的个数、网络下载流所使用的网络带宽等等）。解决的办法就是 **根据资源可用情况限制并发的个数，以匹配合适的并行度**。

接下来，使用缓冲通道来对程序进行改进。创建一个容量为n的缓冲通道来建立一个并发原语，称为 *技术信号量*。对于缓冲通道中的n的空闲槽，每一个代表一个令牌，持有者可以执行。*通过发送一个值到通道来领取令牌，从通道中接收一个值来释放令牌，创建一个新的空闲槽*。这样就可以保证在没有接收操作时，最多同时有n个发送。在这里，因为通道的元素类型并不重要，所以使用空结构体`struct{}`更合适，它占用的空间大小为0。

### crawlv0.2

下面，开始对crawl函数进行重写，使用对令牌的获取的释放来包含对links.Extract()函数的调用。

```go
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
```

第二个问题，这个程序永远不会结束，即是它已经从初始的URL发现了所有的可达链接。为了让程序终止，当任务列表为空且爬取goroutine都结束的时候，应该能够才主循环中退出。

```go
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
```

这个程序中计数器n负责记录任务列表中的任务个数。当有一个条目被发送到任务列表时，就递增变量n，第一次递增发生在发送初始化命令行参数，第二次递增则是发生在每次启动一个新的爬取goroutine时。当主循环从n减到0，则再也没有任务需要完成。

### crawlv0.3

下面这个版本的程序也是一个解决过度并发问题的替代方案。它并没有使用计数信号量，而是通过20个长期存活的爬虫goroutine来调用，确保最多只有20个HTTP请求。

```go
func main() {
  worklist := make(chan []string)   // 可能有重复的URL任务列表
  unseenlinks := make(chan string)  // 去重复的URL任务列表
  go func() { worklist <- os.Args[1:] }()
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
```

第三版的程序逻辑结构图

![程序逻辑](http://on64c9tla.bkt.clouddn.com/Go/images.jpg)
