## goroutine

Go语言中，每一个并发执行的活动称为goroutine。当一个程序启动时，只有一个goroutine调用main函数，称为主goroutine。新的goroutine通过`go`关键字进行创建。

一个简单的例子，计算第45个斐波那契数。主goroutine用于计算斐波那契数，然后通过`go`调用spinner()函数，输出字符串表示程序在运行。

🌰：[斐波那契数](example/spinner.go)

```go
func main() {
  // 创建一个调用spinner()的goroutine
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
```
