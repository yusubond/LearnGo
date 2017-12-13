// Name: test.go
// Func: 测试tempconv包
// Author: subond
// Date: Dec 13, 2017
package main

import (
  "fmt"
  "tempconv"
)

func main() {
  fmt.Println("The AbsoluteZeroC is %v", tempconv.AbsoluteZeroC)
  fmt.Println("The FreezingC is %v", tempconv.FreezingC)
  fmt.Println("The BoilingC is %v", tempconv.BoilingC)
  c := tempconv.F2C(212.0)
  fmt.Println(c.String())
}
