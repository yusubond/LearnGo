package main

import (
  "fmt"
  "flag"
)

func main() {
  /*
  username := flag.String("name", "", "input your name")
  flag.Parse()
  fmt.Println("Hello, ", *username)
  */

  fmt.Println("parsed? = ", flag.Parsed())
  flag.Parse()
  fmt.Println("parsed? = ", flag.Parsed())
  args := flag.Args()
  fmt.Println(args)
  fmt.Println(len(args))
}
