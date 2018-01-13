// Name: main.go
// Func: return the older person
// Author: subond
// Date: 13 Jan, 2018
package main

import "fmt"

type Person struct {
  name string
  age int
}

func Older(p1, p2 Person) (Person, int) {
  if p1.age > p2.age {
    return p1, p1.age - p2.age
  }
  return p2, p2.age - p1.age
}

func main() {
  Tom := Person{"Tom", 12}
  Jack := Person{"Jack", 14}
  John := Person{"John", 23}

  tk, tkage := Older(Tom, Jack)
  th, thage := Older(Tom, John)
  kh, khage := Older(Jack, John)

  fmt.Printf("Of %s and %s, %s is older by %d years\n", Tom.name, Jack.name, tk.name, tkage)
  fmt.Printf("Of %s and %s, %s is older by %d years\n", Tom.name, John.name, th.name, thage)
  fmt.Printf("Of %s and %s, %s is older by %d years\n", Jack.name, John.name, kh.name, khage)
}
