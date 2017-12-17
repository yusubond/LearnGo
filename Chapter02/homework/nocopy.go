// Name: nocopy.go
// Exam: 练习4.5
// Functions: 编写一个就地处理函数，用于去除[]string slice中相邻的重复字符串元素
// Input: []string
// Output: []string
// Eg: input data = "123", "123", "one", "two", "two", "123"
//     output data = "123", "one", "two", "123"
// Author: subond
// Date: Dec 14, 2017
package main

import "fmt"

func main() {
  a := []string{"123", "123", "one", "two", "two", "123"}
  fmt.Println("The Origin data is ", a)
  //fmt.Println("The Origin data is ", a[1:])
  a = nocopy(a)
  fmt.Println("The Reverse data is ", a)
}

func nocopy(data []string) []string {
  k := 0
  for i := 0; i < len(data) - 1; i++ {
    if data[i] != data[i+1] {
      data[i] = data[i+1]
      k = i + 1;
    }
  }
  return data[:k]
}
