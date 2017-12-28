// Name: sort1.go
// Author: subond
// Date: Dec 28, 2017
package main

import (
  "fmt"
  "sort"
)

func main() {
  list := []int{45, 23, 90, 7, 20, 83, 12}
  fmt.Println("The list is ", list)

  sort.Ints(list)
  fmt.Println("The sorted list is ", list)
}
