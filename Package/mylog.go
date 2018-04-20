package main

import (
  "fmt"
  "log"
)

func main() {
  fmt.Println(log.Ldate)
  fmt.Println(log.Ltime)

  fmt.Println(log.Flags())
  fmt.Println(log.Flags())
}
