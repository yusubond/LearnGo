// Name: clock1.go
// Func: 顺序时钟服务器，以每秒一次的频率向客户端发送当前时间
// Author: subond
// Date: Dec 25, 2017
package main

import (
  "io"
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
  defer c.Close()
  for {
    // net.Conn对象满足io.Writer接口，可以直接写入
    // time.Time.Format方法，提供格式化日期和时间信息的方式，它的参数是一个模板
    // 模板为 Mon Jan 2 15:04:05 -0700 MST 2006
    // 参考 [Time.Format](https://golang.org/pkg/time/#Time.Format)
    _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
    if err != nil {
      return
    }
    time.Sleep(1 * time.Second)
  }
}
