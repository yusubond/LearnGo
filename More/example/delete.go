// Name: delete.go
// Func: 练习：使用append函数书写一个delete函数，使其能够满足一下三个情况：
//  1. 删除slice的第一个元素
//  2. 删除slice的最后一个元素
//  3. 删除slice中索引为i的元素，0 < i < len(slice) - 1。
// Author: subond
// Date: 16 Jan, 2018
package main

import "fmt"

func delete(i int, slice []int) []int {
  switch i {
  case 0:
    slice = slice[1:]
  case len(slice) - 1:
    slice = slice[:len(slice) - 1]
  default:
    slice = append(slice[:i], slice[i + 1:]...)
  }
  return slice
}

func main() {
  slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  fmt.Println("The begining slice is: ", slice)
  fmt.Println("Delete the first element")
  slice = delete(0, slice)
  fmt.Println("Now the slice is ", slice)
  fmt.Println("Delete the last element")
  slice = delete(len(slice) - 1, slice)
  fmt.Println("Now the slice is ", slice)
  fmt.Println("Delete the 3rd element")
  slice = delete(3, slice)
  fmt.Println("Now the slice is ", slice)
}
