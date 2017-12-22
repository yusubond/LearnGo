// Name: intset.go
// Exam: 练习6.2
// Func: 实现IntSet的变长方法AddAll(...int)使其可以接收一串整型值作为参数
// Author: subond
// Date: Dec 22, 2017
package main

import (
  "fmt"
  "bytes"
)

// 定义IntSet类型，一个包含非负整数的集合
// 零值代表集合为空
type IntSet struct {
  words []uint64
}

// Add方法，添加非负整数x到集合中
func (s *IntSet) Add(x int) {
  word, bit := x/64, uint(x%64)
  for word >= len(s.words) {
    s.words = append(s.words, 0)
  }
  s.words[word] |= 1 << bit
}

// AddAll方法，允许向集合中添加一串整型值
func (s *IntSet) AddAll(values ...int) {
  for _, val := range values {
    s.Add(val)
  }
}

// String方法，以友好的方式输出集合
func (s *IntSet) String() string {
  var buf bytes.Buffer
  buf.WriteByte('{')
  for i, word := range s.words {
    if word == 0 {
      continue
    }
    for j := 0; j < 64; j++ {
      if word&(1<<uint(j)) != 0 {
        if buf.Len() > len("{") {
          buf.WriteByte(',')
          buf.WriteByte(' ')
        }
        fmt.Fprintf(&buf, "%d", 64*i+j)
      }
    }
  }
  buf.WriteByte('}')
  return buf.String()
}

func main() {
  var x IntSet
  x.Add(1)
  x.Add(9)
  x.Add(65)
  fmt.Println(x.String())             // {1, 9, 65}
  x.AddAll(12, 13, 14)
  fmt.Println(x.String())
}
