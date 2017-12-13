// Name: echo_range.go
// Func: 输出命令的参数即索引，每行一个
// Keys: strconv包, range
// Author: subond
// Date: Dec 10, 2017

package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  for i, arg := range os.Args[1:] {
    s := ""
    s = strconv.Itoa(i) + " " + arg
    //fmt.Print(i)
    fmt.Println(s)
  }
}
