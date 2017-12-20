## 方法

目录：

  * [普通变量作为接受者](#普通变量作为接受者)
  * [指针变量作为接受者](#指针变量作为接受者)

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
