// Name: intset.go
// Func：位向量的使用，用商数x/64作为字的索引，余数x%64作为字内索引
// 例如，集合{1，9，65}
// 存储结果：words[0]中的的第1位和第9位为1
//         words[1]中的第1位为1
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

// Has方法，判断集合中是否包含元素x
func (s *IntSet) Has(x int) bool {
  word, bit := x/64, uint(x%64)
  return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add方法，添加非负整数x到集合中
func (s *IntSet) Add(x int) {
  word, bit := x/64, uint(x%64)
  for word >= len(s.words) {
    s.words = append(s.words, 0)
  }
  s.words[word] |= 1 << bit
}

// UnionWith方法，将集合s和t的并集运算结果存入s中
func (s *IntSet) UnionWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] |= tword
    } else {
      s.words = append(s.words, tword)
    }
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
  var x, y IntSet
  x.Add(1)
  x.Add(9)
  x.Add(65)
  y.Add(9)
  y.Add(68)
  fmt.Println(x.String())             // {1, 9, 65}
  fmt.Println(y.String())             // {9, 68}
  x.UnionWith(&y)
  fmt.Println(x.String())             // {1, 9, 65, 68}
  fmt.Println(x.Has(1), y.Has(65))    // true false
}
