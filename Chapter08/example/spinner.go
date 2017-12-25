// Name: spinner.go
// Func：演示goroutine
// Author: subond
// Date: Dec 25, 2017
package main

import "fmt"
import "time"

func main() {
  go spinner(100 * time.Millisecond)
  const n = 45
  fibN := fib(n)
  fmt.Printf("\nFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
  for {
    for _, r := range `-\|/` {
      fmt.Printf("\r%c", r)
      time.Sleep(delay)
    }
  }
}

func fib(x int) int {
  if x < 2 {
    return x
  }
  return fib(x-1) + fib(x-2)
}
