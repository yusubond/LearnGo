// Name: zero.go
// Func:
// Author: subond
// Date: 16 Jan, 2018
package main
import "fmt"
func main() {
  var slice []int
  fmt.Println("Before calling make")
  if slice==nil { fmt.Println("slice == nil") }
  fmt.Println("len(slice) == ", len(slice))
  fmt.Println("cap(slice) == ", cap(slice))
  // Let's allocate the underlying array:
  fmt.Println("After calling make")
  slice = make([]int, 4)
  fmt.Println("slice == ", slice)
  fmt.Println("len(slice) == ", len(slice))
  fmt.Println("cap(slice) == ", cap(slice))
  // Let's change things:
  fmt.Println("Let's change some of its elements: slice[1], slice[3] = 2, 3")
  slice[1], slice[3] = 2, 3
  fmt.Println("slice == ", slice)
}
