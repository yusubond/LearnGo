## 函数Functions

本质上来讲`函数`就是一个代码块，通过给定输入，返回结果。

目录：

  * [基本语法](#基本语法)
  * [值与引用传递](#值与引用传递)
  * [函数签名](#函数签名)

### 基本语法

一般的函数语法如下：

```go
func funcname(input1 type1, input2 type2) (output1 type1, output2 type2) {
  //some code and processing here
  ...
  //return the results of the function
  return value1, value2
}
```

关键字`func`用来声明一个函数；一个函数可以包含多个输入参数，参数后紧跟着参数的类型，通过`,`隔开；函数需要返回一个或多个结果。

一个简单的Max函数。

```go
func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}
```

举个例子：一个具有两个输出参数的函数。

```go
func SumAndProduct(a, b int) int int {
  return a + b, a * b
}
```

再举个例子：一个具有结果变量的函数。

```go
// A function that returns a bool that is set to true of Sqrt is possible
//and false when not. And the actual square root of a float64
func MySqrt(f float64) (squareroot float64, ok bool){
    if f > 0 {
        squareroot, ok = math.Sqrt(f), true
    } else {
        squareroot, ok = 0, false
}
    return squareroot, ok
}
```

### 值与引用传递

如正面的一个函数所示，当把一个变量传递给函数的时候，函数得到的只是变量的 **一个拷贝**。在函数中改变这个值的大小，并不能真正改变这个值原始的大小。这种称为 **值传递**。

当参数以指针的方式传入函数时，通过取值`*`操作符就可以获得的就是这个变量，这种情况下对变量做出的改变就是有效的，称为 **引用传递**。

举个例子：

```go
// 值传递
func Add1(a int) int {
  return a + 1
}

// 引用传递
func Add2(a *int) int {
  return *a + 1
}
```

### 函数签名

通过前面的介绍，我们可以发现影响一个函数的的因素包括：输入参数、输出参数和函数体。

我们可以重写函数的主体，以不同的方式工作，主程序将继续编译运行，没有问题。

但是我们不能在函数声明中改变输入和输出参数，而仍然在主程序中使用它的旧参数。换句话说，在上面列出的3个元素中，更重要的是：函数期望的输入参数是什么，以及返回的输出参数是什么。

这两个元素就是我们所说的 **函数签名**(`Function's signatures`)，为此我们使用这种形式：

```go
func (input1 type1 [, input2 type2 [, ...]])  (output1 OutputType1 [, output2 OutputType2 [,...]]) {
  ...
}
```

代替一般的函数声明

```go
func function_name (input1 type1 [, input2 type2 [, ...]])  (output1 OutputType1 [, output2 OutputType2 [,...]]) {
  ...
}
```

下面是几个函数签名的例子：

```go
//The signature of a function that takes an int and returns and int
func (int x) x int
//takes two float and returns a bool
func (float32, float32) bool
// Takes a string returns nothing
func (string)
// Takes nothing returns nothing
func()
```
