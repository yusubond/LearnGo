// Name: emptystruct.go
// Func: 空结构体的应用
// Author: subond
// Date: Dec 28, 2017
// Srouce: http://zuozuohao.github.io/2016/06/16/The-empty-struct/
package main

import "fmt"
import "unsafe"

func main() {
  var s string
  var c complex128
  var a int
  fmt.Println(unsafe.Sizeof(s))
  fmt.Println(unsafe.Sizeof(c))
  fmt.Println(unsafe.Sizeof(a))

  var b [3]uint32
  fmt.Println(unsafe.Sizeof(b))
}
