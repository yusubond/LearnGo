// Name: human.go
// Func: 演示函数与接口，fmt.Print与Stringer接口
//       只要一种数据类型，实现了String()方法，就可以直接传递给fmt.Print()函数
// Author: subond
// Date: Dec 28, 2017
package main

import (
  "fmt"
  "strconv"
)

type Human struct {
  name string
  age int
  phone string
}

// Returns a nice string repersenting a human
// With this method, Human implements fmt.Stringer
func (h Human) String() string {
  return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
  Bob := Human{"Bob", 39, "000-7777-XXX"}
  fmt.Println("This Human is : ", Bob)
}
