## 方法

目录：

  * [普通变量作为接受者](#普通变量作为接受者)
  * [指针变量作为接受者](#指针变量作为接受者)
  * [结构体内嵌](#结构体内嵌)
  * [方法变量与表达式](#方法变量与表达式)
  * [封装](#封装)

### 普通变量作为接受者

Go语言中方法是一种特殊类型的函数，在函数名字之前多个一个参数。这个参数表示将这个方法绑定到这个参数对应的类型上，称为这个方法的 *接受者*。下面，写一个与平面几何相关的函数和方法，以示两者的区别。

```go
package geometry

import "math"

// 定义Point结构体
type Point struct {
  X, Y float64
}

// 声明并实现Distance()函数，用来计算两点之间的距离
func Distance(p, q Point) float64 {
  return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// 实现Point类型的方法，用来计算两点之间的距离
func (p Point) Distance(q Point) float64 {
  return math.Hypot(q.X - p.X, q.Y - p.Y)
}
```

上面的程序需要留意的是，两个Distance函数并没有冲突。第一个是一个包级别的函数，称为geometry.Distance()，第二个是一个类型Point的方法，称为Point.Distance。因为每个方法都有一个接受者，所以不同类型中使用相同的方法名是允许的。例如：

```go
// 定义Path类型，表示连接多个点的直线段
type Path []Point

func (path Path) Distance() float64 {
  sum := 0.0
  for i := range path {
    if i > 0 {
      sum += path[i-1].Distance(path[i])
    }
  }
  return sum
}
```

即，类型所拥有的所有方法名都必须是唯一的，而不同类型可以拥有相同的方法名。

### 指针变量作为接受者

当需要更新一个变量的值时，Go语言中采用指针作为方法的接受者，例如：

```go
// 方法名为(*Point).ScaleBy
func (p *Point) ScaleBy(factor float64) {
  p.X += factor
  p.Y += factor
}
```

**命名类(T)和指向它们的指针(*T)是唯一可以出现在接受者声明出的类型**。与此同时，Go语言中为了防止混淆，不允许本身是指针的类型进行方法声明。

总结起来就是，在合法的方法调用表达式中，只有符合以下三种形式的语句才能成立。

1. 实参接受者和形参接受者为同一个类型

```go
p := Point{1, 2}
q := Point{4, 6}
qptr := &q
p.Distance(q)       // 同为Point类型
qptr.ScaleBy(2)     // 同为*Point类型
```

2. 实参接受者是T类型，而形参接受者为*T类型。这个情况下，编译器会进行隐式转换。

```go
var p Point = Point{1, 2}
p.ScaleBy(3)     // 隐式转换(&p)
```

3. 实参接受者为*T类型，而形参接受者为T类型。这个情况下，编译器会进行隐式解引用。

```go
p := Point{1, 2}
q := Point{4, 6}
qptr := &q
qptr.Distance(p)
```

栗子：[mypoint](example/mypoint.go)

### 结构体内嵌

下面定义一种ColorPoint类型

```go
import "image/color"

type Point struct {
  X, Y float64
}

type ColorPoint struct {
  Point
  color.RGBA
}
```

那么这个ColorPoint类型可以同时拥有Point所有的方法，也拥有RGBA所有的方法，以及任何直接在ColorPoint类型中声明的方法。

栗子：[mycolorpoint](example/mycolorpoint.go)

### 方法变量与表达式

通常，我们调用方法时需要指定接受者，比如p.Distance()。但是，Go语言也允许将方法赋值给一个 **方法变量**，它是一个 **函数**，把方法绑定到接受者上。因此，使用方法变量的时候只需要提供实参即可，而不需要提供接受者就可以。

```go
p := Point{1, 2}
q := Point{4, 6}
// 将方法Point.Distance以方法变量的方式赋值到一个方法变量distancefromp
// 这个时候，distancefromp就是一个函数，调用时只提供实参即可
distancefromp := p.Distance
// 调用distancefromp
fmt.Println(distancefromp(q))
```

与方法变量类似的是方法表达式，两者的区别就是：方法表达式写成T.f或(*T.f)，T是类型，它是一种 **函数变量**，把原来方法的接受者作为函数的第一个形参，因此可以像调用其他普通函数一样调用。

```go
p := Point{1, 2}
q := Point{4, 6}
// 将方法Point.Distance以方法表达式的方式赋值到一个变量distance
// 这个时候，distance就是一个函数变量，调用时需要提供包含接受者的实参
distance := Point.Distance
// 调用distance
fmt.Println(distance(q, p))
fmt.Printf("%T\n", distance)  // func(Point, Point) float64
scale := (*Point).ScaleBy
scale(&p, 3)
fmt.Println(p)
fmt.Printf("%T\n", scale)     // func(*Point, float64)
```

栗子1：[pointvar](example/pointvar.go)

栗子2：[pointadd](example/pointadd.go)

习题：用位向量实现集合的运算

参考: [inset](exmaple/intset.go)

### 封装

如果变量或方法是不能通过对象访问的，这称作封装的对象或方法。Go语言中要封装一个对象，**必须使用结构体**。这就是为什么前面的intset类型中虽然只包含一个字段，但是仍然使用的结构体。

```go
type IntSet struct {
  words []uint64
}
```

其次，Go语言中封装的单元是包，而不出类型。*无论是函数内的代码，还是方法内的代码，结构体类型中的字段对同一个包中的所有代码都是可见的*。

封装的好处有三：

第一，因为使用方不能直接修改对象的变量，所以不需要更多的语句来检查变量的值。  
第二，隐藏实现细节可以防止使用方依赖的属性发生变化，使得设计者可以更加灵活的改变API的实现而不破坏兼容性。  
第三，防止使用者肆意地改变对象内的变量。
