// Name: add.go
// Func: parameters by value and by referenc
// Author: subond
// Date: 13 Jan, 2018
package add

func Add1(a int) int {
  return a + 1
}

func Add2(a *int) int {
  *a = *a + 1
  return *a
}
