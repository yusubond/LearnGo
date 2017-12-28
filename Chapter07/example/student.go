// Name: Student.go
// Func: 理解interface type, interface values
// Author: subond
// Date: Dec 28, 2017
package main

import "fmt"

// Human类型
type Human struct {
  name string
  age int
  phone string
}

// Student类型
type Student struct {
  Human
  school string
  load float32
}

// Employee类型
type Employee struct {
  Human
  company string
  money float32
}

// A Human method to say hello
func (h Human) SayHello() {
  fmt.Printf("Hello, I am %s, you can call me on %s\n", h.name, h.phone)
}

// A Human method to sing songs
func (h Human) Sing(lyrics string) {
  fmt.Println("La, la, la...", lyrics)
}

// A Employee method 覆写Human的SayHello
func (e Employee) SayHello() {
  fmt.Printf("Hello, I am %s, I work at %s, you can call me on %s\n", e.name, e.company, e.phone)
}

// A Employee method SpendSalary
func (e *Employee) SpendSalary(amount float32) {
  e.money -= amount
}

// A Student method BorrowMoney
func (s *Student) BorrowMoney(amount float32) {
  s.load += amount
}

// 接口类型
// Human、Student和Employee均实现了接口Men，因为他们都包含Men接口的方法
type Men interface {
  SayHello()
  Sing(lyrics string)
}

func main() {
  mike := Student{Human{"Mike", 25, "13512345670"}, "MIT", 0.0}
  paul := Student{Human{"Paul", 23, "13512345671"}, "Harvard", 100}
  sam := Employee{Human{"Sam", 29, "13512345672"}, "Apple", 1000}
  tom := Employee{Human{"Tom", 35, "13512345673"}, "Google", 5000}

  // 声明一个Men接口类型的变量m
  var m Men

  // m存储Student类型的实例
  m = mike
  fmt.Println("This is Mike, a Student.")
  m.SayHello()
  m.Sing("November rain")

  // m存储Employee类型的实例
  m = sam
  fmt.Println("This is Sam, a Employee")
  m.SayHello()
  m.Sing("Born to be wild")

  // Men类型的slice
  x := make([]Men, 4)
  x[0], x[1], x[2], x[3] = mike, paul, sam, tom
  for _, value := range x {
    value.SayHello()
  }
}
