## 聚合数据类型

目录：

  * [数组](#数组)
  * [slice](#slice)
  * [map](#map)
  * [结构体](#结构体)
  * [习题](#习题)

### 数组

数组是具有 **固定长度** 且拥有零个或多个 **相同数据类型** 元素的序列。内置函数len()可以返回数组中的元素个数，其索引从0开始，到len()-1。

数组的零值是其数据类型的零值。

**数组声明与初始化**

```go
// 声明一个具有三个int对象的数组a
var a [3]int
// 声明一个具有三个int对象的数组b，并将其初始化为1,2,3
var b [3]int = [3]int{1, 2, 3}

var c [3]int = [2]int{4, 5} // 缺省的为0

var d [...]int = [4]int{11, 12, 13, 14}  // ...表示由初始化元素的个数决定，本例中为4

var e [...]int = [...]int{99: -1}  // 按索引赋值，即一共100个元素，除最后一个元素为-1外，其他元素均为0
```

需要记住的是，数组长度是数组类型的一部分，而且必须使常量表达式，即`[3]int`和`[4]int`是不同的数组类型。

多维数组

```go
// 声明并初始化一个2行4列的二维数组
double_array := [2][4]int {[4]int{1,2,3,4}, [4]int{5,6,7,8}
// 也可以用下面的这种方式
double_array := [2][4]int {{1,2,3,4}, {5,6,7,8}}
```

### slice

slice表示一个拥有相同数据类型元素的可变长度的序列，其底层数据结构为数组，它有三个属性，**指针、长度和容量**。内置函数len()和cap()可以返回slice的长度和容量。

因为slice包含指向数组元素的指针，因此将slice传递给函数的时候，可以在函数内部修改底层数组元素。

```go
var s []int = []int{1, 2, 3}
a := b[i:j]   // 当b是一个数组时，操作符[i:j]可以创建并返回一个新的slice
// make([]Type, len, cap)
c = make([]int, 3)
d = make([]int, 3, 5)  // 内置make函数创建slice。
```

**slice的零值是nil**。值为nil的slice没有对应的底层数组，其长度和容量都是零。

内置make()函数创建slice，其实质就是创建一个无名数组，并返回它的的一个slice。

举个🌰：[zero.go](example/slice_functions/zero.go)

**为什么使用make而不是new创建slice?**

调用函数`new(T)`的意思是分配存储类型的T的存储空间并返回一个指针指向这个空间，因此`new(T)`的返回是`*T`。但我们使用`new([]int)`创建一个int类型的切片的时候，它的返回是`*[]int`并不是我们所期望的结果。

所以，我们才会希望有这样一种函数：

1. 它能够返回一个实际的数据结构，而不是一个指针。就像上面的例子，我们希望的是[]int，而不是*[]int。

2. 能够分配底层的数组用于存储切片的元素。

因此，`make`函数不仅仅是从主机获得足够的空间用来存储数据结构，而是分配相应的数据结构和底层数组。

内置append()函数可以将元素追加到slice中。

```go
var runes []rune
for _, r = range "Hello World!" {
  runes = append(runes, r)
}
```

**slice的工作原理**

因为slice的底层数据结构是数组，因此下面AppendInt()函数可以帮助我们很好地理解其工作原理。

```go
func AppendInt(x []int, y int) []int {
  var z []int
  zlen = len(x) + 1
  if zlen <= cap(x) {
    // 当x的容量仍可以装下新元素时
    z = x[:zlen]
  } else {
    // 当x的容量不能装下新的元素时，需要进行扩容，这里采用2倍扩容策略
    zcap := zlen
    if zcap < 2 * len(x) {
      zcap = 2 * len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z, x)
  }
  z[len(x)] = y
  return z
}
```

由于内置append()函数使用了比我们举例所用的AppendInt()函数更加复杂的扩容策略，因此在实际调用append()函数的时候，我们并不清楚哪一次调用会导致一次新的内存分配，也就无法断定原始的slice和返回的slice结果是否指向同一个底层数据。

**slice与函数**

举个例子：找出最大的一个数。

```go
func Max(slice []int) int {
  max := slice[0]
  for i := 0; i < len(slice); i++ {
    if slice[i] > max {
      max = slice[i]
    }
  }
  return max
}
```
一个例子：[PrintByteSlice](example/slice_functions/reference.go)

说明slice只是数组或其他切片的一个引用。

### map

map是散列表的引用，其类型是map[k]v，map中所有的键都有相同的数据类型，同时所有的值也都有相同的数据类型。内置函数make()可以用来创建一个map。

```go
map[keytype] valuetype
```

这种类型与Python中的 *dict* 和 Ruby中的*hash*相似。

```go
// 声明一个从string到int映射的map
var numbers map[string]int

// 用make函数声明并初始化
age1 := make(map[string]int)
age2 := make(map[string]int){
  "alice": 31,
  "charlie": 25,
}

age1["alice"] = 34

// 删除一个元素
delete(age1, "alice")
```

map的零值也是nil，与slice一样，即没有引用任何散列表。

通过下标的方式访问map中的元素，可以输出两个值，第一个值就是key所对应的值，第二个值是一个布尔值，表示该元素是否存在。

**range函数**

Go语言为我们提供了一个range函数，可用于迭代slice，array，map等数据结构。

### 结构体

结构体是将零个或多个任意类型的命名变量组合在一起的聚合数据类型。其定义如下：

```go
// 定义了一个名为Employee的结构体数据类型，其包含ID, Name, Address, Salary四个成员
type Employee struct {
  ID      int
  Name    string
  Address string
  Salary  int
}
```

如果一个结构体的成员变量名称是首字母大写，那么这个变量是可导出的，如果是小写字母，则不可以。通过点号的方式可以访问结构体成员。

结构体的值可以通过结构体字母量来设置，有两种方式：一种是直接设置，如下所示。

```go
type Point struct {X, Y int}
p := Point{1, 2}
```

还有一种就是通过指定部分或全部成员变量的名称来初始化结构体的值，需要注意的是当成员变量是不可导出的时候，不能使用这个方法。如下所示。

```go
package p1
type T struct {a, b int}

package p2

import "p1"
var v = p1.T{a: 1, b: 2}  // 编译错误，无法引用a, b
```

**结构体嵌套机制**

当多个结构体存在重复变量的时候，Go语言中的结构体嵌套机制可以减少代码的重复性，如下所示。

```go
// 定义Point结构体，表示一个点
type Point struct {
  X, Y int
}
// 定义Circle结构体，表示一个圆
type Circle struct {
  Center Point
  Radius int
}
// 定义Wheel结构体，表示一个轮子
type Wheel struct {
  Circle Circle
  Spokes int
}
```

当然，为了更好的书写和阅读代码，Go语言允许定义不带名称的结构体成员，只要指定类型即可。这样的成员称为 **匿名成员**。因此，上面的Circle和Wheel结构体可以写成下面的格式：

```go
type Circle struct {
  Point     // 匿名成员
  Radius int
}
type Wheel struct {
  Circle    // 匿名成员
  Spokes int
}
```

正是有了这样的嵌套机制，对于结构体变量的访问也更加方便。

```go
var w Wheel
w.X = 1         // 等价于x.Circle.Point.X = 1
w.Y = 3         // 等价于x.Circle.Point.Y = 3
w.Radius = 5    // 等价于x.Circle. Radius = 5
w.Spokes = 20
```

### 区别与联系

数组和结构体都是聚合数据类型，它们的值由内存中的一组变量构成。数组元素具有相同的数据类型，而结构体中的元素数据类型可以不同。数组和结构体的长度都是固定的，而slice和map都是动态数据结构。

### 习题

1. [older person](example/older10/older10.go)
