// Name: du1.go
// Func: 并发目录遍历
// Author: subond
// Date: 5 Jan, 2018
package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
)

// walkDir递归遍历以dir为根目录的整个文件树
// 并在fileSizes上发送每个已经找到的文件的大小
func walkDir(dir string, fileSizes chan<- int64) {
  for _, entry := range dirents(dir) {
    if entry.IsDir() {
      subdir := filepath.Join(dir, entry.Name())
      walkDir(subdir, fileSizes)
    } else {
      fileSizes <- entry.Size()
    }
  }
}

// dirents返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
  entries, err := ioutil.ReadDir(dir)
  if err != nil {
    fmt.Fprintf(os.Stderr, "du1: %v\n", err)
    return nil
  }
  return entries
}

func main() {
  // 确定初始目录
  flag.Parse()
  roots := flag.Args()
  if len(roots) == 0 {
    roots = []string{"."}
  }

  // 遍历文件树
  fileSizes := make(chan int64)
  go func() {
    for _, root := range roots {
      walkDir(root, fileSizes)
    }
    close(fileSizes)
  }()

  // 输出结果
  var nfiles, nbytes int64
  for size := range fileSizes {
    nfiles++
    nbytes += size
  }
  printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
  fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
