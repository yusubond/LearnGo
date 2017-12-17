// Name: reverse2.go
// Exam: 练习4.3
// Functions: 就地反转整型数组中的元素，使用指针而不是silce
// Input: []int
// Output: NONE
// Author: subond
// Date: Dec 14, 2017
package main

import "fmt"

const SIZE = 6

func main() {
  a := [SIZE]int{1, 2, 3, 4, 5, 6}
  fmt.Println("The Origin data is ", a)
  reverse(&a)
  fmt.Println("The Reverse data is ", a)
}

func reverse(data *[SIZE]int) {
  for i, j := 0, len(data) - 1; i < j; i, j = i + 1, j - 1 {
    data[i], data[j] = data[j], data[i]
  }
}
