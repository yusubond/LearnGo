// Name: thumbnail5.go
// Func: 并行处理生成缩略图问题
// Author: subond
// Date: Dec 26, 2017
package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "gopl.io/ch8/thumbnail"
)

func main() {
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    thumb, err := thumbnail.ImageFile(input.Text())
    if err != nil {
      log.Print(err)
      continue
    }
    fmt.Println(thumb)
  }
  if err := input.Err(); err != nil {
    log.Fatal(err)
  }
}
