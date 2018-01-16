// Name: max2.go
// Func: 找出一个int slice中最大值
// Author: subond
// Date: 16 Jan, 2018
package main

import "fmt"

func max(slice []int) int {
  max := slice[0]
  for _, value := range slice {
    if value > max {
      max = value
    }
  }
  return max
}

func main() {
  var s1 = []int{1, 2, 3, 4, 7}
  fmt.Printf("The max number of the slice is %d\n", max(s1))
}
