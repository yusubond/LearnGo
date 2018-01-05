## 聊天服务器

聊天服务器允许几个用户之间互相广播文本消息。这个程序里有4个goroutine。主goroutine和广播goroutine(broadcaster)，每一个连接里面有一个连接处理goroutine(handleConn)和一个客户写入goroutine(clientWriter)。

```go
// 对外发送消息的通道
type client chan<- string

// 所有接收的客户消息
var (
  entering = make(chan client)
  leaving  = make(chan client)
  messages = make(chan string)
)

func broadcaster() {
  // 所有连接的客户端
  clients := make(map[client]bool)
  for {
    select {
    case msg := <-messages:
      // 把所有接收的消息广播给所有的用户
      // 发送消息通道
      for cli := range clients {
        cli <- msg
      }
    // 客户到来
    case cli := <-entering:
      clients[cli] = true
    // 客户离开
    case cli := <-leaving:
      delete(clients, cli)
      close(cli)
    }
  }
}

func handleConn(conn net.Conn) {
  // 对外发送客户消息的通道
  ch := make(chan string)
  go clientWriter(conn, ch)

  who := conn.RemoteAddr().String()
  ch <- "You are " + who
  messages <- who + " has arrived"
  entering <- ch

  input := bufio.NewScanner(conn)
  for input.Scan() {
    messages <- who + ": " + input.Text()
  }

  leaving <- ch
  messages <- who + " has left"
  conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
  for msg := range ch {
    fmt.Fprintln(conn, msg)
  }
}

func main() {
  listener, err := net.Listen("tcp", "localhost:8080")
  if err != nil {
    log.Fatal(err)
  }

  go broadcaster()
  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Print(err)
      continue
    }
    go handleConn(conn)
  }
}
```
