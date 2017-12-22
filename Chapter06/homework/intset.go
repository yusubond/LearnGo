// Name: intset.go
// Exam: 练习6.1
// Func: 实现IntSet的以下方法
//       func (*IntSet) Len() int       // 返回集合中元素个数
//       func (*IntSet) Remove(x int)   // 从集合中移除元素x
//       func (*IntSet) Clear()         // 清空集合元素
//       func (*IntSet) Copy() *IntSet  // 返回集合的副本
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

// Remove方法，从集合中删除元素x
func (s *IntSet) Remove(x int) {
  word, bit := x/64, uint(x%64)
  for i, _ := range s.words {
    if i == word {
      s.words[i] &^= 1 << bit
      break
    }
  }
}

// Clear方法，清空所有元素
func (s *IntSet) Clear() {
  for i, _ := range s.words {
    s.words[i] &= 0
  }
}

// Len方法，返回集合中的元素个数
func (s *IntSet) Len() int {
  var count = 0
  for _, word := range s.words {
    for j := 0; j < 64; j++ {
      if word & (1 << uint(j)) != 0 {
        count++
      }
    }
  }
  return count
}

// Copy方法，返回集合的副本
func (s *IntSet) Copy() *IntSet {
  var res IntSet
  for _, word := range s.words {
    res.words = append(res.words, word)
  }
  return &res
}

// Has方法，判断集合中是否包含元素x
func (s *IntSet) Has(x int) bool {
  word, bit := x/64, uint(x%64)
  return word < len(s.words) && s.words[word]&(1<<bit) != 0
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
  fmt.Println(x.Len())                // 3
  x.Remove(9)
  fmt.Println(x.String())             // {1, 65}
  x.Remove(1)
  fmt.Println(x.String())             // {65}
  x.Remove(5)
  fmt.Println(x.String())             // {65}
  x.Clear()
  fmt.Println(x.String())             // {}
  fmt.Println(y.String())             // {9, 68}
  zp := y.Copy()
  fmt.Println(zp.String())
  x.UnionWith(&y)
  fmt.Println(x.String())             // {1, 9, 65, 68}
  fmt.Println(x.Has(1), y.Has(65))    // true false
}
