// Name: mycolorpoint.go
// Func: 说明方法可在类型中的传递性
// Author: subond
// Date: Dec 20, 2017

package main

import "math"
import "fmt"
import "image/color"

// 定义Point类型
type Point struct {
  X, Y float64
}

// 实现Point类型的方法，用来计算两点之间的距离
func (p Point) Distance(q Point) float64 {
  return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
  p.X += factor
  p.Y += factor
}

// 定义ColorPoint类型
type ColorPoint struct {
  Point
  color.RGBA
}

func main() {
  red := color.RGBA{255, 0, 0, 255}
  blue := color.RGBA{0, 0, 255, 255}
  p := ColorPoint{Point{1, 2}, red}
  q := ColorPoint{Point{4, 6}, blue}
  // 虽然ColorPoint类型拥有Point类型的所有方法，但是其传递的参数仍为Point类型中指定的参数类型
  // 即Distance(q Point)
  fmt.Println(p.Distance(q.Point))
  p.ScaleBy(3)
  fmt.Println(p.Distance(q.Point))
}
