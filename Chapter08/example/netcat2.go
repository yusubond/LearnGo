// Name: netcat.go
// Func: 顺序时钟客户端
// Author: subond
// Date: Dec 25, 2017
package main

import (
  "io"
  "log"
  "net"
  "os"
)

func main() {
  conn, err := net.Dial("tcp", "localhost:8080")
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()
  // 创建一个goroutine，读取服务器的回复到标准输出
  go mustCopy(os.Stdout, conn)
  //主goroutine 从标准输入读取，并发送到服务器
  mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
  if _, err := io.Copy(dst, src); err != nil {
    log.Fatal(err)
  }
}
