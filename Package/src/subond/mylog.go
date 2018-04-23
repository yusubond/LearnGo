package subond

import (
  "fmt"
  "log"
)

func MyLog() {
  fmt.Println(log.Ldate)
  fmt.Println(log.Ltime)
  //  log.Fatal("there is something error, so exit.")
  //  log.Fatalf("there is %d errors, so exit.", 5)
  fmt.Println(log.Flags())
  // log.Panic("something")
  fmt.Println(log.Prefix())
}
