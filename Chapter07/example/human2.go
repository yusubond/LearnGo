// Name: human2.go
// Func: 实现sort接口，包含Len() Less() Swap()方法
// Author: subond
// Date: Dec 28, 2017
package main

import (
  "fmt"
  "sort"
  "strconv"
)

type Human struct {
  name string
  age int
  phone string
}

func (h Human) String() string {
  return "(name: " + h.name + " - age: " + strconv.Itoa(h.age)  + " years)"
}

type HumanGroup []Human

func (hg HumanGroup) Len() int {
  return len(hg)
}

func (hg HumanGroup) Less(i, j int) bool {
  if hg[i].age < hg[j].age {
    return true
  }
  return false
}

func (hg HumanGroup) Swap(i, j int) {
  hg[i], hg[j] = hg[j], hg[i]
}

func main() {
  mygroup := HumanGroup{
    Human{name:"Mike", age:23},
    Human{name:"Paul", age:39},
    Human{name:"Toms", age:12},
    Human{name:"Jeck", age:25},
    Human{name:"Vean", age:19},
  }

  fmt.Println("The unsorted group is:")
  for _, value := range mygroup {
    fmt.Println(value)
  }

  sort.Sort(mygroup)

  fmt.Println("The sorted group is:")
  for _, value := range mygroup {
    fmt.Println(value)
  }
}
