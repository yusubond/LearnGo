// Name: older10.go
// Func:
// Author: subond
// Date: 16 Jan, 2018
package main

import "fmt"

type person struct {
  name string
  age int
}

// 返回10个人中年龄最大的一个
func Older10(person [10]person) person {
  older := person[0]
  for i := 1; i < 10; i++ {
    if person[i].age > older.age {
      older = person[i]
    }
  }
  return older
}

func main() {
  var array [10]person

  // 没有指定的元素，默认为person{"", 0}
  array[1] = person{"Paul", 23}
  array[2] = person{"Jim", 24}
  array[3] = person{"Sam", 84}
  array[4] = person{"Rob", 54}
  array[8] = person{"Karl", 19}

  /*
  第二种初始化方式
  array := [10]person{
    person{"", 0},
    person{"Paul", 23},
    person{"Jim", 24},
    person{"Sam", 84},
    person{"Rob", 54},
    person{"", 0},
    person{"", 0},
    person{"", 0},
    person{"Karl", 10},
    person{"", 0}}
  */
  older := Older10(array)

  fmt.Printf("The older of the group is %s\n", older.name)
}


double_array := [2][4]int {[4]int{1,2,3,4}, [4]int{5,6,7,8}
