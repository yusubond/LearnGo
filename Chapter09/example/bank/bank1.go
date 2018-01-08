// Name: bank1.go
// Func: 通过teller的监控goroutine限制balance变量
// Author: subond
// Date: 8 Jan, 2018
package bank

var deposits = make(chan int)   // 发送存款额
var balances = make(chan int)   // 接收余额

func Deposit(amount int) { deposits <- amount }
func Balance() int { return <-balances }

func teller() {
  var balance int
  for {
    select {
    case amount := <-deposits:
      balance += amount
    case balances <- balance:
    }
  }
}

func init() {
  go teller()
}
