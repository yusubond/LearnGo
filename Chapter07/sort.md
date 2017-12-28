### sort包与interface

sort包用来对int、float、以及string数据类型排序，也用到了接口类型。先从一个小🌰开始。

```go
package main

import (
  "fmt"
  "sort"
)

func main() {
  list := []int{45, 23, 90, 7, 20, 83, 12}
  fmt.Println("The list is ", list)

  sort.Ints(list)
  fmt.Println("The sorted list is ", list)
}
```

事实上，sort包定义了一个包含三个方法的接口类型：

```go
type Interface interface {
  // Len is the number of elements in the collection.
  Len() int
  // Less returns whether the element with index i shoud sort
  // before the element with index j
  Less(i, j int) bool
  // Swap swaps the elements with indexes i and j
  Swap(i, j int)
}
```

其实，在sort包的文档中可以看到：

> type, typically a collection, that satisfies sort. Interface can be sorted by the routines in this package. The methods require that the elements of the collection be enumerated by an integer index.

所以，给一个给定的slice排序，只要实现Len(),Less(),Swap()三个方法就可以！

```go
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
```

其输出结果为：

```shell
The unsorted group is:
(name: Mike - age: 23 years)
(name: Paul - age: 39 years)
(name: Toms - age: 12 years)
(name: Jeck - age: 25 years)
(name: Vean - age: 19 years)
The sorted group is:
(name: Toms - age: 12 years)
(name: Vean - age: 19 years)
(name: Mike - age: 23 years)
(name: Jeck - age: 25 years)
(name: Paul - age: 39 years)
```

由此可见，我们并没有实现HumanGroup的排序函数，只是实现了三个方法(Len(), Less(), Swap())，而这就是sort.Sort()函数所需的全部信息。

其实质是，Sort包的排序函数接受任意类型的参数，只要他实现了Sort接口类型。

🌰：[HumanGroup](example/human2.go)
