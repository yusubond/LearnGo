// Name: intset.go
// Exam: 练习6.3
// Func: 实现IntSet的交集、差集和对称差集
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

// 并集：UnionWith方法，将集合s和t的并集运算结果存入s中
func (s *IntSet) UnionWith(t *IntSet) *IntSet {
  var res IntSet
  for _, sword := range s.words {
    res.words = append(res.words, sword)
  }
  for i, tword := range t.words {
    if i < len(s.words) {
      res.words[i] |= tword
    } else {
      res.words = append(res.words, tword)
    }
  }
  return &res
}

// 交集：InterSectionWith方法，返回集合s和t的交集结果
// 既属于集合s也属于集合t的元素所构成的集合，即集合s ∩ 集合t
func (s *IntSet) InterSectionWith(t *IntSet) *IntSet {
  var res IntSet
  for i, sword := range s.words {
    for j, tword := range t.words {
      if i == j {
        res.words = append(res.words, (sword & tword))
      }
    }
  }
  return &res
}

// 差集：DifferenceWith方法，返回集合s和t的差集结果
// 所有属于集合s但并不属于集合t所构成的集合，即集合s - 集合t
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
  var res IntSet
  for _, sword := range s.words {
    res.words = append(res.words, sword)
  }
  for i, _ := range s.words {
    for j, tword := range t.words {
      if i == j {
        res.words[i] &^= tword
      }
    }
  }
  return &res
}

// 对称差集：所有只在一个集合存在的元素，即(s ∪ t) - (s ∩ t)
func (s *IntSet) SyDifferenceWith(t *IntSet) *IntSet {
  stu := s.UnionWith(t)
  sti := s.InterSectionWith(t)
  res := stu.DifferenceWith(sti)
  return res
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
  x.AddAll(12, 13, 14, 65, 67, 68)
  y.AddAll(4, 5, 14, 65, 90, 150)
  fmt.Printf("集合x为:\n%s\n", x.String())             // {1, 9, 65}
  fmt.Printf("集合y为:\n%s\n", y.String())             // {1, 9, 65}
  xyi := x.InterSectionWith(&y)
  fmt.Printf("集合x与y的交集为:\n%s\n",xyi.String())
  xyu := x.UnionWith(&y)
  fmt.Printf("集合x与y的并集为:\n%s\n",xyu.String())
  xyd := x.DifferenceWith(&y)
  fmt.Printf("集合x与y的差集为:\n%s\n",xyd.String())
  xysd := x.SyDifferenceWith(&y)
  fmt.Printf("集合x与y的对称差集为:\n%s\n",xysd.String())
}
