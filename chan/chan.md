## chan

### 什么是chan

在golang中`chan`，即channel，译为通道。每一个通道是一个具体类型的导管，称作通道的`元素类型`。

```go
ch := make(chan int)  // ch类型为chan int
```
通道是一个使用`make`创建的数据结构的引用。*当复制或者作为参数传递到一个函数的时，复制的是引用，这样的调用者和被调用者都引用同一份数据结构。*和其他引用类型一样，通道的零值是`nil`。

通道支持三种操作：发送(send)、接收(receive)、和关闭(close)。
```go
send: 一个goroutine传输一个值到另一个执行接收表达式的goroutine
receive: 一个goroutine接收另一个goroutine发送的值
close: 
```


### chan底层数据结构

```go
type hchan struct {
	qcount   uint           // total data in the queue// 队列中数据的大小
	dataqsiz uint           // size of the circular queue// chan底层数据结构是循环队列，其值为循环队列大小
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32  // 表示通道是否关闭
	elemtype *_type // element type  // 通道的元素类型
	sendx    uint   // send index        // 发送索引
	recvx    uint   // receive index      // 接收索引
	recvq    waitq  // list of recv waiters  // 接受者goroutine的队列，这是一个双向链表
	sendq    waitq  // list of send waiters // 发送者goroutine的队列 ，这是一个双向链表

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex // goroutine锁
}
// 接受者和发送者的底层结构
type waitq struct {
	first *sudog
	last  *sudog
}

// sudog represents a g in a wait list, such as for sending/receiving
// on a channel.
//
// sudog is necessary because the g ↔ synchronization object relation
// is many-to-many. A g can be on many wait lists, so there may be
// many sudogs for one g; and many gs may be waiting on the same
// synchronization object, so there may be many sudogs for one object.
//
// sudogs are allocated from a special pool. Use acquireSudog and
// releaseSudog to allocate and free them.
// sudog用来表示等待队列中的一个goroutine
//  这是一个随时可以使用的中间变量或元素
type sudog struct {
	// The following fields are protected by the hchan.lock of the
	// channel this sudog is blocking on. shrinkstack depends on
	// this for sudogs involved in channel ops.

	g *g

	// isSelect indicates g is participating in a select, so
	// g.selectDone must be CAS'd to win the wake-up race.
	isSelect bool
	next     *sudog  // 后继节点指针
	prev     *sudog  // 前继节点指针 这表示双向链表
	elem     unsafe.Pointer // data element (may point to stack)

	// The following fields are never accessed concurrently.
	// For channels, waitlink is only accessed by g.
	// For semaphores, all fields (including the ones above)
	// are only accessed when holding a semaRoot lock.

	acquiretime int64
	releasetime int64
	ticket      uint32
	parent      *sudog // semaRoot binary tree
	waitlink    *sudog // g.waiting list or semaRoot
	waittail    *sudog // semaRoot
	c           *hchan // channel
}
```

### chan创建


### chan接收


### chan发送