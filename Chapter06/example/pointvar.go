// Name: pointvar.go
// Func: 说明方法变量与方法表达式的区别
//       方法变量：一个特殊(绑定了接受者)的函数，调用的时候只需要提供实参即可
//       方法表达式：一个特殊(将接受者作为第一个形参)的函数变量，调用的时候按顺序提供实参
// Author: subond
// Date: Dec 21, 2017

package main

import "math"
import "fmt"

// 定义Point结构体
type Point struct {
  X, Y float64
}

// 实现Point类型的Distance方法，用来计算两点之间的距离
func (p Point) Distance(q Point) float64 {
  return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// 实现Point类型的ScaleBy方法，用于节点的伸缩
func (p *Point) ScaleBy(factor float64) {
  p.X += factor
  p.Y += factor
}

func main() {
  p := Point{1, 2}
  q := Point{4, 6}
  // 方法变量
  distancefromp := p.Distance
  fmt.Println(distancefromp(q))   // 5
  scalep := p.ScaleBy
  scalep(3)
  fmt.Println(p)                  // {4, 5}
  // 方法表达式
  distance := Point.Distance
  fmt.Println(distance(q, p))     // 1
  fmt.Printf("%T\n", distance)    // func(Point, Point) float64
  scale := (*Point).ScaleBy
  scale(&p, -3)
  fmt.Println(p)                  // {1, 2}
  fmt.Printf("%T\n", scale)       // func(*Point, float64)
}
