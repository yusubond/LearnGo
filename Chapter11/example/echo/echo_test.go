// Name: echo_test.go
// Func: echo测试代码
// Author: subond
// Date: 12 Jan, 2018
package main

import (
  "bytes"
  "fmt"
  "testing"
)

func TestEcho(t *testing.T) {
  var tests = []struct {
    newline bool
    sep     string
    args    []string
    want    string
  }{
    {true, "", []string{}, "\n"},
    {false, "", []string{}, ""},
    {true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
    {true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
    {false, ":", []string{"1", "2", "3"}, "1:2:3"},
    {true, ",", []string{"a", "b", "c"}, "a b c\n"},
  }
  for _, test := range tests {
    descr := fmt.Sprintf("echo(%v, %q, %q)", test.newline, test.sep, test.args)
    out = new(bytes.Buffer)   // 捕获的输出
    if err := echo(test.newline, test.sep, test.args); err != nil {
      t.Errorf("%s failed: %v", descr, err)
      continue
    }
    got := out.(*bytes.Buffer).String()
    if got != test.want {
      t.Errorf("%s = %q, want %q", descr, got, test.want)
    }
  }
}
