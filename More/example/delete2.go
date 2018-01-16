// Name: delete2.go
// Func: 更加一般形式的删除函数
// Author: subond
// Date: 16 Jan, 2018
package main

import "fmt"

func delete(i int, slice []int) []int {
  if slice == nil {
    return slice
  }
  if i > len(slice) {
    return slice
  } else {
    slice = append(slice[:i], slice[i + 1:]...)
    return slice
  }
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
  fmt.Println()

  var slice2 []int
  fmt.Println("The begining slice2 is ", slice2)
  if slice2 == nil { fmt.Println("slice2 == nil") }
  slice2 = delete(1, slice2)
  fmt.Println("Now the slice2 is ", slice2)
}
