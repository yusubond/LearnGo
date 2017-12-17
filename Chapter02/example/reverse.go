// Name: reverse.go
// Functions: 就地反转整型slice中的元素
// Input: []int
// Output: NONE
// Author: subond
// Date: Dec 14, 2017
package main

import "fmt"

func main() {
  a := []int{1, 2, 3, 4, 5, 6}
  fmt.Println("The Origin data is ", a)
  reverse(a)
  fmt.Println("The Reverse data is ", a)
}

func reverse(data []int) {
  for i, j := 0, len(data) - 1; i < j; i, j = i + 1, j - 1 {
    data[i], data[j] = data[j], data[i]
  }
}
