// Name: du2.go
// Func: 并发目录遍历，-v可显示进度条
//       为每一个walkDir的调用创建一个新的goroutine
// Author: subond
// Date: 5 Jan, 2018
package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "sync"
  "path/filepath"
  "time"
)

// walkDir递归遍历以dir为根目录的整个文件树
// 并在fileSizes上发送每个已经找到的文件的大小
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
  defer n.Done()
  for _, entry := range dirents(dir) {
    if entry.IsDir() {
      n.Add(1)
      subdir := filepath.Join(dir, entry.Name())
      go walkDir(subdir, n, fileSizes)
    } else {
      fileSizes <- entry.Size()
    }
  }
}

// sema是一个用于限制目录并发数的计数信号量
var sema = make(chan struct{}, 20)

// dirents返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
  sema <- struct{}{}          // 获取令牌
  defer func() { <-sema }()   // 释放令牌
  entries, err := ioutil.ReadDir(dir)
  if err != nil {
    fmt.Fprintf(os.Stderr, "du1: %v\n", err)
    return nil
  }
  return entries
}

var verbose = flag.Bool("v", false, "show verbose progerss messages")

func main() {
  // 确定初始目录
  flag.Parse()
  roots := flag.Args()
  if len(roots) == 0 {
    roots = []string{"."}
  }

  // 遍历文件树
  fileSizes := make(chan int64)
  var n sync.WaitGroup
  for _, root := range roots {
    n.Add(1)
    go walkDir(root, &n, fileSizes)
  }
  go func() {
    n.Wait()
    close(fileSizes)
  }()

  // 输出结果
  var tick <-chan time.Time
  if *verbose {
    tick = time.Tick(500 * time.Millisecond)
  }
  var nfiles, nbytes int64
loop:
  for {
    select {
    case size, ok := <-fileSizes:
      if !ok {
        break loop
      }
      nfiles++
      nbytes += size
    case <-tick:
      printDiskUsage(nfiles, nbytes)
    }
  }
  printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
  fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
