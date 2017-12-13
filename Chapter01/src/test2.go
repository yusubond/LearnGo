// Name: popcount_test.go
// Func: 测试popcount
// Author: subond
// Date: Dec 13, 2017
package main

import "fmt"
import "popcount"

func main() {
  var x uint64 = 34
  fmt.Printf("The 1 bits in %d is %v", x, popcount.Popcount(x))
}
