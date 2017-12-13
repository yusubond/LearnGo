// Name: dup2.go
// Func: 读取指定文件，输出其中出现次数超过1行的内容
// Author: subond
// Date: Dec 10, 2017
package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func main() {
  counts := make(map[string]int)
  for _, filename := range os.Args[1:] {
    data, err := ioutil.ReadFile(filename)
    // 读取文件的错误处理
    if err != nil {
      fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
      continue
    }
    for _, line := range strings.Split(string(data), "\n") {
      counts[line]++
    }
  }
  for line, n := range counts {
    if n > 1{
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
