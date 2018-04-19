package main

import (
  "fmt"
  "os"
  //"os/signal"
  "log"
)

func main() {
  /*
  文件打开
  file, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err)
  }

  data := make([]byte, 500)
  count, err := file.Read(data)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("read %d bytes: %q\n", count, data[:count])
  */

  /*
  信号
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)
  s := <-c
  fmt.Println("Got signal:", s)
  */

  /*
  目录切换及当前目录
  fmt.Println(os.Getpid())
  fmt.Println(os.Getppid())
  err := os.Chdir("/Users/subond/Documents/MyLife/LearnGo")
  if err != nil {
    log.Fatal(err)
  }
  */
  pwd, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(pwd)

  hostname, err := os.Hostname()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(hostname)

  demofile, err := os.Create("demo.txt")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(demofile.Name())

  testfile, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(testfile.Name())
  data := make([]byte, 200)
  _, err = testfile.Read(data)
  if err != nil {
    log.Fatal(err)
  }
  size, err := demofile.Write(data)
  fmt.Printf("write %d bytes to %s file.\n", size, demofile.Name())
}
