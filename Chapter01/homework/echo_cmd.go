// Name: echo_cmd.go
// Func: 输出命令的名字
// Author: subond
// Date: Dec 10, 2017

package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println(os.Args[0])
}
