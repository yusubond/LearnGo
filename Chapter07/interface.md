## 接口

目录

  * [接口类型](#接口类型)
  * [接口变量](#接口变量)
  * [空接口类型](#空接口类型)
  * [接口参数与函数](#接口参数与函数)
  * [接口实现](接口实现)
  * 应用[sort包与interface](sort.md)

像前面讲的的基本类型、聚合类型以及函数、方法等都属于 **具体类型**。具体类型指定了它所含数据的精确布局，还暴露了基于这个精确布局的内部操作。如果我们拿到一个具体类型的值，那么我们就可以精确地知道 *它是什么* 以及 *它能干什么*。

接口类型是一个 **抽象类型**，它提供的仅仅是一种方法而已。当我们拿到一个接口类型的值，我们无从知道 *它是什么*，你能知道的仅仅是 *它能做什么*，更精确地讲，仅仅是 *它提供了哪些方法*。

当提到接口的时候，我们需要明白两件事情：1.接口是方法的一个集合，2.接口也是一种数据类型**。`An interface is two things: it is a set of methods, but it is also a type.`

### 接口类型

接口就是一组方法签名的集合。**一个接口类型定义了一套方法，如果一个具体类型要实现该接口，那么必须实现接口类型定义中的所有方法**。我们使用接口来识别一个对象能够进行的操作。

假设，Student实现了SayHello、Sing和BorrowMoney方法，Employee实现了SayHello、Sing和SpendSalary方法。那么，这些方法的集合就是Student和Employee满足的接口类型。比如，Student和Employee都满足包含SayHello、Sing方法签名的接口，但Student不满足包含SayHello、Sing和SpendSalary方法签名的接口，因为Student没有SpendSalary的方法。

有了上面的解释，我们就可以定义相应的接口。

```go
type Human struct {
  name, phone string
  age int
}

type Student struct {
  Human
  school string
  load float32
}

type Employee struct {
  Human
  company string
  money float32
}

func (h *Human) SayHello() {
  fmt.Printf("Hello, I am %s, you can call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing() {
  fmt.Println("La, la, la...")
}

// 覆写Human的SayHello
func (e *Employee) SayHello() {
  fmt.Printf("Hello, I am %s, I work at %s, you can call me on %s", e.name, e.company, e.phone)
}

func (e *Employee) SpendSalary(amount float32) {
  e.money -= amount
}

func (s *Student) BorrowMoney(amount float32) {
  s.load += amount
}

// 接口类型
type Men interface {
  SayHello()
  Sing()
}

type YoungChap interface {
  SayHello()
  Sing()
  BorrowMoney(amount float32)
}

type ElderlyGent interface {
  SayHello()
  Sing()
  SpendSalary(amount float32)
}
```

由此，可以看到，一个接口可以被任意数量的类型满足，同时，一个类型也实现任意数量的接口。这里Student和Employee都实现了Men接口，而且，Student也实现了YoungChap接口，Employee实现了ElderlyGent接口。

最值得注意的是，**每个类型都实现了一个空接口interface{}**。

### 接口变量

既然接口是一种数据类型，那么它一定一个对应的值。如果你声明了一个接口变量，那么这个变量能够存储 **任何实现该接口的对象类型**。对于上面的例子，如果你声明了一个Men接口类型的变量m，那么这个变量m可以存储Student和Employee类型的对象，当然还有Human类型的对象，因为他们都实现了Men接口类型。

如果m可以存储不同数据类型的值，那么就可以实现一个Men切片(slice)，该切片包含不同数据类型的实例。

🌰:[Men接口类型及变量](example/student.go)

由此可以发现：接口类型是一组抽象方法的集合，其本身并不实现方法或精确地描述数据结构和方法的实现方式。另外，需要提及的是，这些数据类型(Student和Employee类型)也没有提及接口的信息，方法的实现部分也没有提及接口的信息。

同样的，接口类型也不关心到底是什么数据类型实现了自身，因为Men接口也没有提及Student和Employee类型的信息。

所以，接口类型的本质就是 **如果一种数据类型实现了自身的方法集，那么该接口类型变量就可以引用这种数据类型的值**。

### 空接口类型

空接口类型interface{}，一个方法都不包含，所以数据类型都实现了它。虽然空接口类型在描述一个对象实例时明显不足，但是它却可以存储任何数据类型，这就可以极大地发挥它的用武之地。

```go
// 声明一个空接口变量
var a interface{}
var i int = 5
s := "Hello world"
// 因为空接口类型可以存储任何数据类型的实例，所以，以下表达式都是合法的
a = i
a = s
```

如果一个函数的参数包含一个空接口类型interface{}，那么它的意思是 **可以接收任何数据类型**。如果一个函数返回一个空接口类型，那么函数的意思是：“我也不确定返回什么，你只要知道我一定返回一个值就好了”。

### 接口类型与函数

知道一个接口类型如何存储满足它的数据类型实例，以及如何存储不同数据类型实例的集合，我们就可以让函数来接受满足特定接口类型的数据类型实例。

我们知道fmt.Print()函数是一个可变参数的函数，可以接收任意数量的参数，而且既可以是string类型，也可以是int，甚至float类型。这其中的奥秘就与interface有关。

实际上在fmt包中，可以看到如下的接口声明：

```go
// The Stringer interface in fmt package
type Stringer interface {
  String() string
}
```

意思是，任何数据类型，只有实现Stringer接口（具体一点，就是String()方法），就可以传递给fmt.Print函数，并打印出该数据类型String()方法返回的值。

```go
// Returns a nice string repersenting a human
// With this method, Human implements fmt.Stringer
func (h Human) String() string {
  return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}
```

### 接口实现

如果一个类型实现了一个接口要求的所有方法，那么这个类型实现了这个接口。

正如前面所讲，接口类型是一组抽象方法的集合，其本身并不实现方法或精确地描述数据结构和方法的实现方式。另外，需要提及的是，这些数据类型(Student和Employee类型)也没有提及接口的信息，方法的实现部分也没有提及接口的信息。

### [sort包与interface](sort.md)
