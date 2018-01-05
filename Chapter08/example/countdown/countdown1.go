// Name: countdown1.go
// Func: select多路复用
// Author: subond
// Date: 5 Jan, 2018
package main

import "fmt"
import "time"

func main() {
  fmt.Println("Commencing countdown")
  tick := time.Tick(1 * time.Second)
  for countdown := 10; countdown > 0; countdown-- {
    <-tick
  }
  launch()
}

func launch() {
  fmt.Println("Lift Off!")
}
