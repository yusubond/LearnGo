## 位向量

目录：

  * [基本方法](#基本方法)
  * [集合运算](#集合运算)

在数据流分析领域，集合元素都是小的非负整数，集合拥有多个元素，但是集合的操作多数为求并集和交集，这种情况下 *位向量* 就是理想的数据结构。

位向量使用一个无符号整型值的slice，每一位代表集合中的一个元素。如果设置第i位的元素，则认为集合包含元素i。

```go
// IntSet是一个包含非负整型值的集合
// 零值表示集合为空
type IntSet struct {
  words []uint64
}
```

例如，集合{1，9，65}，用位向量表示就是：words[0]中的的第1位和第9位为1，words[1]中的第1位为1。

### 基本方法

下面，演示一下集合的集中方法。

  * Add()：将一个元素x添加到集合中

```go
// Add方法，即添加元素x到集合中
func (s *IntSet) Add(x int) {
  word, bit := x/64, uint64(x%64)
  // 将元素x之前的位置填满，因此用for，而不是if
  for word >= len(s.words) {
    s.words = append(s.words, 0)
  }
  s.words[word] |= 1 << bit
}
```

  * AddAll()：将若干个元素添加到集合中

```go
// AddAll，即添加若干个元素到集合中
func (s *IntSet) AddAll(values ...int) {
  for _, val := range values {
    word, bit := val/64, uint64(val%64)
    for word >= len(s.words) {
      s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
  }
}
```

  * Has：即判断集合中是否包含某个元素

```go
// Has方法，判断集合中是否包含某个元素
func (s *IntSet) Has(x int) bool {
  word, bit := x/64, uint64(x%64)
  return word < len(s.words) && (s.words[word] & (1 << bit) != 0)
}
```

  * Remove：即从集合中移除某个元素

```go
// Remove方法，从集合中删除元素x
func (s *IntSet) Remove(x int) {
  word, bit := x/64, uint(x%64)
  for i, _ := range s.words {
    if i == word {
      // &^按位置零
      s.words[i] &^= 1 << bit
      break
    }
  }
}
```
  * Clear：即清空集合

```go
// Clear方法，清空所有元素
func (s *IntSet) Clear() {
  for i, _ := range s.words {
    s.words[i] &= 0
  }
}
```

  * Len：即返回集合中的元素个数

```go
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
```

  * Copy：即返回一个集合的副本

```go
// Copy方法，返回集合的副本
func (s *IntSet) Copy() *IntSet {
  var res IntSet
  for _, word := range s.words {
    res.words = append(res.words, word)
  }
  return &res
}
```

  * String方法，以"{1, 2, 3}"的形式返回集合元素

```go
// String方法，以"{1, 2, 3}"形式返回集合元素
func (s *IntSet) String() string {
  var buf bytes.Buffer
  buf.WriteByte('{')
  for i, word := range s.words {
    if word == 0 {
      continue
    }
    for j := 0; j < 64; j++ {
      if word & (1 << uint(j)) != 0 {
        if buf.Len() > len("{") {
          buf.WriteByte(',')
          buf.WriteByte(' ')
        }
        fmt.Fprintf(&buf, "%d", 64 * i + j)
      }
    }
  }
  buf.WriteByte('}')
  return buf.String()
}
```

### 集合运算

  * 交集运算，即返回两个集合的交集集合

```go
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
```

  * 并集运算，返回两个集合的并集结果

```go
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
```

  * 差集运算，返回集合s与集合t的差集结果

```go
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
```

  * 对称差集，返回两个集合的对称差集结果

```go
// 对称差集：所有只在一个集合存在的元素，即(s ∪ t) - (s ∩ t)
func (s *IntSet) SyDifferenceWith(t *IntSet) *IntSet {
  stu := s.UnionWith(t)
  sti := s.InterSectionWith(t)
  res := stu.DifferenceWith(sti)
  return res
}
```
