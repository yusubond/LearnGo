// Name: conv.go
// Func: 包tempcov负责摄氏温度与华氏温度间的转换
// File: tempcov.go conv.go
// Author: subond
// Date: Dec 13, 2017
package tempconv

func C2F(c Celsius) Fahrenheit {
  return Fahrenheit(c * 9 / 5 +32)
}
func F2C(f Fahrenheit) Celsius {
  return Celsius((f - 32) * 5 / 9)
}
