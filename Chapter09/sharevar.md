## 使用共享变量实现并发

### 竞态

串行程序中各个步骤的执行顺序由程序逻辑决定，每个goroutine内部的各个步骤也是顺序执行的。当一个程序中包含两个goroutine或多个goroutine时，如果一个goroutine中的事件x和另一个goroutine中的事件y的执行顺序是不确定的，那么我们称这两个事件是 *并发* 的。

如果一个函数在并发调用时仍然能够正确工作，那么这个函数就是 *并发安全* 。

当多个goroutine按某些交错顺序执行时，程序不能给出正确结果，称为 *竞态*。

数据竞态发生于 两个goroutine **并发读写**同一个变量，并且 **至少其中一个是写入时**。

三种避免数据竞态的方法：

1. 不要修改变量。

那么从不修改的数据结构以及不可变数据结构本质上是并发安全的，也不需要做任何同步。

2. 避免从多个goroutine访问同一个变量。

在[并发的Web爬虫](Chapter08/concurrenceweb.md)中主goroutine是唯一一个能够访问seen map的goroutine；[聊天服务器](Chapter08/chat.md)中的broadcaster goroutine是唯一一个能访问client map的goroutine。这些变量都 **限制在单个goroutine内部**。

> 不要通过共享内存来通信，而应该通过通信来共享内存。

由于其他goroutine无法直接访问相关变量，因此它们就必须使用通道来向受限goroutine发送查询请求或更新变量。

```go
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
```

3. 允许多个goroutine访问同一个变量，但在 **同一时间只有一个goroutine可以访问**，即互斥机制。

### 互斥锁

一个容量为1的通道可以保证在同一个时间最多有一个goroutine能访问共享变量。一个计数上限为1的信号量称为 **二进制信号量**。

对上面的例子以二进制信号量的方式进行实现：

```go
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
```

sync包中有一个单独的Mutex类型来支持这种模式，它的Lock()方法用于获取令牌(称为上锁)，Unlock()方法用于释放令牌(称为解锁)。

下面用sync中的Mutex类型重写bank程序：

```go
import "sync"

var (
  mu sync.Mutex           // 保护balance
  balance int
)

func Deposit(amount int) {
  mu.Lock()
  balance = balance + amount
  mu.Unlock()
}

func Balance() int {
  mu.Lock()
  b := balance
  mu.Unlock()
  return b
}
```

**一个goroutine在每次访问共享变量之前，必须先调用互斥量的Lock方法来获取一个互斥锁**。如果其他goroutine已经取走了互斥锁，那么操作会一直阻塞到其他goroutine调用Unlock之后。

互斥量用来 **保护** 共享变量。

在Lock和Unlock之间的代码，可以自由地读取和修改共享变量，这一区域称为 **临界区域**。

记住：一个goroutine使用完成后必须释放互斥锁，否则其他goroutine将无法获取互斥锁。

每个函数在开始时申请一个互斥锁，在结束时再释放掉，通过这种方式来确保共享变量不会被并发访问。这个函数、互斥锁、变量的组合方式称为 **监控** 模式。

为了确保函数的所有分支(包括错误分支)都必须Lock和Unlock成对执行，才能确保互斥的正确使用。而Go语言中的defer语句，**通过延迟执行Unlock** 就可以把临界区域隐式扩展到当前函数的结尾，避免了必须在一个或多个远离Lock的位置插入一条Unlock语句。

```go
func Balance() int {
  mu.Lock()
  defer mu.Unlock()
  return balance
}
```
