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
  mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
  if _, err := io.Copy(dst, src); err != nil {
    log.Fatal(err)
  }
}
