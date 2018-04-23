package subond

import (
  "fmt"
  "os"
  "log"
)

func OpenFile() {
  file, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Open the file test.txt %v\n", file)  // *File指针

  // 读数据
  data := make([]byte, 100)
  count, err := file.Read(data)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("the file is %d bytes\nthe content is : %s\n", count, string(data))
}

func WordDictory() {
  // 获取 切换目录
  wd, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("the work dictory is ", wd)
  os.Chdir("/Users/subond/Documents/MyLife/LearnGo")
  newwd, err := os.Getwd()
  fmt.Printf("switch to the %s dictory\n", newwd)
  os.Chdir(wd)
  nwd, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Now the wd is ", nwd)
}

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
  /*
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
*/
