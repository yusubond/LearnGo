// Name: mypoint.go
// Func: 实现二维平面中的线段计算
// Author: subond
// Date: Dec 20, 2017

package main

import "math"
import "fmt"

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

func (p *Point) ScaleBy(factor float64) {
  p.X += factor
  p.Y += factor
}

func main() {
  var p Point = Point{1, 2}
  var q Point = Point{4, 6}
  qptr := &q
  fmt.Println(Distance(p, q))
  fmt.Println(p.Distance(q))
  p.ScaleBy(3)
  fmt.Println(p.Distance(q))
  fmt.Println(qptr.Distance(p))
}
