// Name: dup.go
// Func: 由标准输入读取元素，输出出现次数大于1的内容
// Keys: bufio包
// Author: subond
// Date: Dec 10, 2017
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  // 内置函数make用来新建map类型
  // map存储一个键值对集合，本例中键的类型为string，值为int
  counts := make(map[string]int)
  imput := bufio.NewScanner(os.Stdin)
  for imput.Scan() {
    counts[imput.Text()]++
  }
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
