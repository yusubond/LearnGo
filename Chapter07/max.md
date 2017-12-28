## Max与Older

其实，`max(s []int) int`和`Older(p []Person) Person`也有类似的功能。实际上，实现一个slice的最大值就是在做一件事情：迭代处理和比较。来看下面这个例子。

```go
// 定义Person类型
type Person struct {
  name string
  age int
}

// 定义三种slice类型
type IntSlice []int
type Float32Slice []float32
type PersonSlice []Person

// 定义MaxInterface接口类型
type MaxInterface interface {
  // Len 返回切片中元素的个数
  Len() int
  // Get 返回索引为i的元素
  // 注意：以空接口类型interface{}作为返回值，说明可以返回任意数据类型
  Get(i int) interface{}
  // Bigger 返回一个布尔值，代表索引i的元素比索引j的元素大
  Bigger(i, j int) bool
}

// 分别实现三种slice类型的Len方法
func (i IntSlice) Len() int { return len(i) }
func (f Float32Slice) Len() int { return len(f) }
func (p PersonSlice) Len() int { return len(p) }

// 分别实现三个slice类型的Get方法
func (x IntSlice) Get(i int) interface{} { return x[i] }
func (x Float32Slice) Get(i int) interface{} { return x[i] }
func (x PersonSlice) Get(i int) interface{} {return x[i] }

// 分别实现三种slice类型的Bigger方法
func (x IntSlice) Bigger(i, j int) bool {
  if x[i] > x[j] {
    return true
  }
  return false
}

func (x Float32Slice) Bigger(i, j int) bool {
  if x[i] > x[j] {
    return true
  }
  return false
}

func (x PersonSlice) Bigger(i, j int) bool {
  if x[i].age > x[j].age {
    return true
  }
  return false
}

// Person类型的String()方法，方便优雅地输出信息
func (p Person) String() string {
  return "(name: " + p.name + " age: " + strconv.Itoa(p.age) + " years)"
}

/*
Returns a bool and a value
- The bool is set to true if there is a MAX in the collection
- The value is set to the MAX value or nil, if the bool is false
 */
func Max(data MaxInterface) (ok bool, max interface{}) {
  if data.Len() == 0 {
    return false, nil
  }
  if data.Len() == 1 {
    return true, data.Get(1)
  }
  max = data.Get(0)
  m := 0
  for i := 1; i < data.Len(); i++ {
    if data.Bigger(i, m) {
      max = data.Get(i)
      m = i
    }
  }
  return true, max
}

func main() {
  islice := IntSlice{9, 18, 23, 6, 14, 10, 0, 21}
  fslice := Float32Slice{3.4, 12.9, 0.12, 6.01, 31.23, 10.34}
  pslice := PersonSlice{
    Person{"Mike", 23},
    Person{"Paul", 39},
    Person{"Toms", 12},
    Person{"Jack", 25},
    Person{"Vean", 19},
  }

  _, m := Max(islice)
  fmt.Println("The biggest integer in islice is :", m)

  _, m = Max(fslice)
  fmt.Println("The biggest float in fslice is :", m)

  _, m = Max(pslice)
  fmt.Println("The oldest person in the pslice is:", m)
}
```

几个关键的点：

1. Max()函数以一个接口类型变量`data MaxInterface`作为形参，其意思是：只要一种具体数据类型实现了MaxInterface接口，那么这个接口变量均可以存储该数据类型的值，因此可以作为Max()函数的参数。

这是利用了 **接口变量可以存储任意实现该接口的数据类型的值** 的原则。

2. Get()方法以空接口类型`interface{}`作为方法的返回值，意思是：可以返回任意数据类型。

因为Get()方法的形参只要索引值，并不知道传入的slice的具体数据类型，它必须保证根据索引就可以返回元素类型。而空接口类型正是可以包含所有数据类型，作为返回值最合适不过了。

这是利用了 **空接口类型interface{}可以充当任意数据类型** 的原则。

3. 接口变量作为函数的形参。
