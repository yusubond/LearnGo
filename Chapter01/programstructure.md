## 基本程序结构

目录：

  * [helloworld](#helloworld)
  * [名称-声明-变量](#名称-声明-变量)
  * [指针](#指针)
  * [类型声明](#类型声明)
  * [包和文件](#包和文件)

一切语言都是从基本组件构建而来，变量存储值，简单表达式通过运算合并成大的语言块；基本类型通过数组和结构体进行聚合；表达式通过if、for等控制语言来决定执行顺序；语言被组织成函数用于隔离和复用；函数被组织成源文件和包。

### helloworld

从Hello World程序开始，简单介绍go程序的结构组成。下面，go语言的hello world程序源码。

```go
// Name: helloworld.go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

Go语言使用包来组织代码，其关键字为package。package指明当前文件所属的包，一个包由一个或多个.go源文件组成。如上所示，`package main `指明了helloworld.go文件属于main包。需要注意的是，**main包是一个特殊的包，它用来定义一个独立的可执行程序**。同样，main包里面的main函数也特殊，它是程序开始执行的地方。例如，

```go
tempcov
  |-tempcov.go
  |-conv.go
```
包名为tempcov，包含两个源文件tempcov.go和conv.go。

接下来是import声明，其 **必须** 跟在package声明之后。import声明指出了组成程度的函数、变量、常量、类型声明。当需要导入多个包时，使用`()`列表。比如，

```go
import (
  "fmt"
  "math"
)
```
凡是被导入的包，在后序的程序里面可以直接使用。如果使用了未导入的包，会报错`undefined xxx`错误；如果导入了某个包而没有使用，则会报错`imported and not used: xxx`。

另外，gofmt工具可以格式化go文件，一个良好的程序员应该养成使用gofmt的工具，让自己的代码更加标准和格式化。

最后，`fmt.Println()`是fmt包中的Println函数。值得提醒的是，go语言对名称的大小写字母敏感。**一个实体第一个字母的大小写决定其可见性是否跨包。包名总是由小写字母组成**。如fmt包中的Println函数的第一个字母是大写，意味着在包外可用，因此，helloworld.go文件可直接使用。

### 名称-声明-变量

对于，Go语言名称需要知道以下几件事情：

1. Go语言中的变量名称对大小写敏感。  
2. 实体第一个字母的大小写决定其可见性是否跨包。如果以大写字母开头，是可导出的，意味着其在包外是可见和可访问的。  
3. 包名本身总是以小写字母组成。  
4. 名称采用“驼峰式”的命名风格。

声明包括四个组要声明：变量(var)、常量(const)、类型(type)和函数(func)。每一个文件以package声明开头，表明文件属于哪个包；紧跟着的是import声明，然后是包级别的类型、变量、常量、函数的声明。

```go
package main

import "fmt"

const SIZE = 20

func main() {
  var m int = 10
  fmt.Println("%d", m * SIZE)
}
```

常量SIZE是一个包级别的声明，属于main包。m是属于main函数的局部变量。**包级别的实体不仅对包含其声明的源文件可见，对同一个包里的所有源文件也可见**。

关键字`var`声明一个具体类型的变量，赋予名字，设置初值。其中，初值设置可以省略，这种情况下编译器自动为其分配该类型的零值。这种零值保障机制很好的保证了go语言中不存在未初始化的变量。

```g0
var name type = expression
```

还有一种 **短变量声明** 的可选形式，它使用`name := expression`的格式，name的类型由expression决定。这种方式更加短小、灵活，非常好用。注意，`:=`表示声明，而不是赋值。还有一个重要的细节：短变量声明不需要声明所有左侧的变量。如果一些变量已经在同一个作用域中声明，那么对于这些变量，短声明的行为等同于赋值。例如，

```
in, err := os.Open(filename)
out, err := os.Create(file) // 这个短声明中仅声明了out变量，而对err变量进行赋值
```

短变量声明最少声明一个新变量，否则，编译无法通过。

### 指针

变量是存储值的地方，而指针的值是一个变量的地址。**每一个聚合类型变量（结构体的成员和数组中的元素）的组成都是变量，因此也有一个地址**。取地址使用操作符为&，取值的操作符为*。

指针的零值为nil。让函数返回局部变量的地址是非常安全的。下面来看new函数。

内置的new()函数可以创建变量。表达式new(T)，表示创建一个未命名的T类型变量，初始化为T类型的零值，并 **返回其地址**。

```go
p := new(int) // *int类型的p，指向未命名的int变量
fmt.Print(*p) // 输出"0"
*p = 2        // 把未命名的int变量设置为2
fmt.Print(*p) // 输出"2"
```

每次调用new()函数返回一个具有唯一地址的不同变量。

### 类型声明

关键字`tpye`定义一个新的命名类型，它和已有类型使用相同的底层类型。

```go
tpye name underlying-type
```
例如，不同计量单位间的温度值转换。

```go
package tempcov

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC     Celsius = 0
  BoilingC      Celsius = 100
)

func C2F(c Celsius) Fahrenheit {
  return Fahrenheit(c * 9 / 5 +32)
}

func F2C(f Fahrenheit) Celsius {
  return Celsius((f - 32) * 5 / 9)
}
```

包tempcov分别定义两个类型：Celsius(摄氏温度)和Fahrenheit(华氏温度)，均使用相同的底层类型float64。尽管它们使用了相同的底层类型，但是它们仍然是不同的类型，不能直接转换，需要进行Celsius(t)和Fahrenheit(t)显示转换。**Celsius(t)和Fahrenheit(t)是类型转换，而不是函数调用**。

Go语言中对于每个类型T，都有一个对应的类型转换操作T(x)，将值x转换成类型T。

### 包和文件

每一个包给它的声明提供了独立的命名空间。package声明前面紧挨着文档注释对整个包进行描述。以温度值转换为例，我们来编写tempcov包。

```shell
# 文件目录
tempconv
  |-tempconv.go
  |-conv.go
```

```go
// Name: tempconv.go
// 将类型、常量、方法的声明均放在tempcov.go文件中
// tempcov包负责摄氏温度和华氏温度间的转换
package tempconv

import "fmt"

// 声明两个类型，Celsius和Fahrenheit
type Celsius float64
type Fahrenheit float64

// 声明三个包级别的常量
const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC     Celsius = 0
  BoilingC      Celsius = 100
)

// 声明String()方法
func (c Celsius) String() string { return fmt.Sprintf("%g˚C", c)}
func (f Fahrenheit) String() string { return fmt.Sprintf("%g˚F", f)}
```

```go
// Name: conv.go
package tempconv

func C2F(c Celsius) Fahrenheit {
  return Fahrenheit(c * 9 / 5 +32)
}
func F2C(f Fahrenheit) Celsius {
  return Celsius((f - 32) * 5 / 9)
}
```

然后，测试包tempconv的文件可以这样写，

```go
// Name: test.go
// 测试包tempconv
package main

import (
  "fmt"
  "tempconv"
)

func main() {
  fmt.Println("The AbsoluteZeroC is %v", tempconv.AbsoluteZeroC)
  c := tempconv.F2C(212.0)
  fmt.Println(c.String())
}
```

包的初始化从初始化包级别的变量开始。对于包级别的每一个变量，生命周期从其值被初始化开始。任何文件都可以包含任意数量的声明，如init函数：

```go
func init() { /* ... */ }
```

这个init函数不能被调用和被引用，在程序启动的时候，init函数按照它们生命的顺序自动执行。

例如，包popcount中的[init()](src/popcount/popcount.go)函数。
