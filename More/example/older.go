// Name: person.go
// Func: 以变量函数的方式重写older函数
// Author: subond
// Date: 16 Jan, 2018
package main

import "fmt"

type person struct {
  name string
  age int
}

func Older(people ...person) (bool, person) {
  if len(people) == 0 {
    return false, person{}
  }
  older := people[0]
  for _, value := range people {
    if value.age > older.age {
      older = value
    }
  }
  return true, older
}

func main() {
  var (
    ok bool
    older person
  )
  paul := person{"Paul", 23};
  jim := person{"Jim", 24};
  sam := person{"Sam", 84};
  rob := person{"Rob", 54};
  karl := person{"Karl", 19};
  // Who is older? Paul or Jim?
  _, older = Older(paul, jim) //notice how we used the blank identifier
  fmt.Println("The older of Paul and Jim is: ", older.name)
  // Who is older? Paul, Jim or Sam?
  _, older = Older(paul, jim, sam)
  fmt.Println("The older of Paul, Jim and Sam is: ", older.name)
  // Who is older? Paul, Jim , Sam or Rob?
  _, older = Older(paul, jim, sam, rob)
  fmt.Println("The older of Paul, Jim, Sam and Rob is: ", older.name)
  // Who is older in a group containing only Karl?
  _, older = Older(karl)
  fmt.Println("When Karl is alone in a group, the older is: ", older.name)
  // Is there an older person in an empty group?
  ok, older = Older() //this time we use the boolean variable ok
  if !ok {
      fmt.Println("In an empty group there is no older person")
  }
}
