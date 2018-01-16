## 函数进阶

目录：

  * [变长函数](#变长函数)
  * [递归函数](#递归函数)
  * [defer声明](#defer声明)

### 变长函数

所谓变量函数是指可以接收同一类型的不定数量的输出参数，其声明如下：

```go
func FuncName(args ...InputType) (output1 OutType1 [, output2 OutType2 [, ...]])
```
举个例子：用变长函数重写Older函数。

```go
func Older(people ...person) (bool, person) {
  if len(people) == 0 {
    return false, person{}
  }
  older := people[0]
  for _, value := range people {
    if value.age > older.age {
      older = value
    }
  }
  return true, older
}
```

内置函数`append`就是一种变长函数。

练习：使用append函数书写一个delete函数，使其能够满足一下三个情况：

1. 删除slice的第一个元素

2. 删除slice的最后一个元素

3. 删除slice中索引为i的元素，0 < i < len(slice) - 1。

```go
func delete(i int, slice []int) []int {
  switch i {
  case 0:
    slice = slice[1:]
  case len(slice) - 1:
    slice = slice[:len(slice) - 1]
  default:
    slice = append(slice[:i], slice[i + 1:]...)
  }
  return slice
}
```

### 递归函数

递归函数是指函数可以调用自己，属于算法一部分，这里不做详细介绍

### defer声明


参考资料：go-book
