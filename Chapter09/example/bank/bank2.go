// Name: bank2.go
// Func: 使用互斥机制保护balance变量
// Author: subond
// Date: 8 Jan, 2018
package bank

var (
  sema = make(chan struct{}, 1)  // 包含balance的二进制信号量
  balance int
)

func Deposit(amount int) {
  sema <- struct{}{}      // 发送操作，获取令牌
  balance = balance + amount
  <-sema                  // 接收操作，释放令牌
}

func Balance() int {
  sema <- struct{}{}      // 发送操作，获取令牌
  b := balance
  <-sema                  // 接收操作，释放令牌
  return b
}
