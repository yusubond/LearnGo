// Name: pointadd.go
// Func: 利用方法表达式，创建函数量，并调用
// Author: subond
// Date: Dec 22, 2017

package main

import "fmt"

// 定义Point结构体
type Point struct {
  X, Y float64
}

type Path []Point

// Point类型的Add方法
// 类型为func(p, q Point) Point
func (p Point) Add(q Point) Point { return Point{q.X + p.X, p.Y + q.Y} }

// Point类型的Sub方法
// 类型为func(p, q Point) Point
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

// Path的TranslateBy方法
func (path Path) TranslateBy(offset Point, add bool) {
  // 声明op函数变量的类型，与Point.Add和Point.Sub方法类型一样
  var op func(p, q Point) Point
  if add {
    op = Point.Add
  } else {
    op = Point.Sub
  }
  for i := range path {
    // 调用Add或Sub方法
    path[i] = op(path[i], offset)
  }
}

func main() {
  p1 := Point{1, 2}
  p2 := Point{3, 3}
  x := []Point{{1, 2}, {2, 4}, {3, 8}}
  translate := Path.TranslateBy
  translate(x, p1, true)
  fmt.Println(x)            // [{2, 4}, {3, 6}, {4, 10}]
  translate(x, p2, false)
  fmt.Println(x)            // [{-1, 1}, {0, 3}, {1, 7}]
}
