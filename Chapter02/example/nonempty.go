// Name: nonempty.go
// Functions: 去除字符串中的空字符串
// Input: slice []string
// Output: slice []string
// Author: subond
// Time: Dec 14, 2017
package main

import "fmt"

func main() {
  var s = []string{"123", "", "one", "", "two", "three"}
  fmt.Printf("No Empty is %s\n", s)
  s = nonempty(s)
  fmt.Printf("No Empty is %s\n", s)
}

func nonempty(strings []string) []string {
  i := 0
  for _, s := range strings {
    if s != "" {
      strings[i] = s
      i++
    }
  }
  return strings[:i]
}
