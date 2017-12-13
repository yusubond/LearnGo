// Name: tempconv.go
// Func: 包tempcov负责摄氏温度与华氏温度间的转换
// File: tempconv.go conv.go
// Author: subond
// Date: Dec 13, 2017
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
