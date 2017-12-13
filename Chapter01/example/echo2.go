// Name: echo2.go
// Author: subond
// Date: Dec 10, 2017
package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println(strings.Join(os.Args[1:], " "))
}
