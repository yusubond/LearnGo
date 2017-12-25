// Name: reverb2.go
// Func: 并发回声服务器
// Author: subond
// Date: Dec 25, 2017
package main

import (
  "bufio"
  "fmt"
  "strings"
  "log"
  "net"
  "time"
)

func main() {
  // Listen()函数，创建一个net.listener对象，用来监听端口
  listener, err := net.Listen("tcp", "localhost:8080")
  if err != nil {
    log.Fatal(err)
  }
  for {
    // Accept方法，返回一个net.Conn对象，代表一个连接
    conn, err := listener.Accept()
    if err != nil {
      log.Print(err)
      continue
    }
    //handleConn(conn)        // 一次处理一个连接
    go handleConn(conn)     // 并发处理多个连接

  }
}

func handleConn(c net.Conn) {
  input := bufio.NewScanner(c)
  for input.Scan() {
    go echo(c, input.Text(), 1 * time.Second)
  }
  defer c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
  fmt.Fprintln(c, "\t", strings.ToUpper(shout))
  time.Sleep(delay)
  fmt.Fprintln(c, "\t", shout)
  time.Sleep(delay)
  fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
