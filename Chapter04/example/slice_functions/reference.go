// Name: reference.go
// Func: slice with reference
// Author: subond
// Date: 16 Jan, 2018
package main

import "fmt"

func PrintByteSlice(name string, slice []byte) {
  fmt.Printf("%s is : [", name)
  for i := 0; i < len(slice) - 1; i++ {
    fmt.Printf("%q,", slice[i])
  }
  fmt.Printf("%q]\n", slice[len(slice) - 1])
}

func main() {
  // Declare an array of 10 bytes.
  A := [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
  //declare some slices.
  slice1 := A[3:7] // Slice1 == {'d', 'e', 'f', 'g'}
  slice2 := A[5:] // Slice2 == {'f', 'g', 'h', 'i', 'j'}
  slice3 := slice1[:2] // Slice3 == {'d', 'e'}
  // Let's print the current content of A and the slices.
  fmt.Println("First content of A and the slices")
  PrintByteSlice("A", A[:])
  PrintByteSlice("slice1", slice1)
  PrintByteSlice("slice2", slice2)
  PrintByteSlice("slice3", slice3)

  A[4] = 'E'
  fmt.Println("\nContent of A and the slices, after changing 'e' to 'E' in array A")
  PrintByteSlice("A", A[:])
  PrintByteSlice("slice1", slice1)
  PrintByteSlice("slice2", slice2)
  PrintByteSlice("slice3", slice3)
}
