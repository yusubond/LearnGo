## 基本控制流

### if

```go
// 直接比较
if x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is less than 10")

// 短变量声明
if x := computed_value(); x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is less than 10")

// 多个
if integer == 3 {
    fmt.Println("The integer is equal to 3")
} else if integer < 3 {
    fmt.Println("The integer is less than 3")
} else {
    fmt.Println("The integer is greater than 3"
```

### for

基本语法：

```go
for expression1; expression2; expression3{
   ...
```

几个例子：

```go
// 第一种
sum := 0
for index:=0; index < 10 ; index++ {
  sum += index
}

// 第二种
sum := 1
for ; sum < 1000; {
    sum += sum
}
// 第三种
for {
    fmt.Println("I loop for ever!")
}
```

### break和continue

`break`跳出当前循环体；`continue`跳出本次循环迭代过程，进入下一次迭代。

### switch

基本语法：

```go
switch sExpr {
  case expr1:
      some instructions
  case expr2:
      some other instructions
  case expr3:
      some other instructions
  default:
      other code
```

几个例子：

```go
// 第一种
i := 10
switch i {
  case 1:
      fmt.Println("i is equal to 1")
  case 2, 3, 4:
      fmt.Println("i is equal to 2, 3 or 4")
  case 10:
      fmt.Println("i is equal to 10")
  default:
      fmt.Println("All I know is that i is an integer")
}

// 第二种
index := 10
switch {
  case index < 10:
      fmt.Println("The index is less than 10")
  case index > 10, index < 0:
      fmt.Println("The index is either bigger than 10 or less than 0")
  case index == 10:
      fmt.Println("The index is equal to 10")
  default:
      fmt.Println("This won't be printed anyway")
}

// 第三种
integer := 6
switch integer {
case 4:
    fmt.Println("The integer was <= 4")
    fallthrough
case 5:
    fmt.Println("The integer was <= 5")
    fallthrough
case 6:
    fmt.Println("The integer was <= 6")
    fallthrough
case 7:
    fmt.Println("The integer was <= 7")
    fallthrough
case 8:
    fmt.Println("The integer was <= 8")
    fallthrough
default:
    fmt.Println("default case")
}
```
