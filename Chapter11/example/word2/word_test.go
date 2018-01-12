// Name: word_test.go
// Func: package word的测试文件
// Author: subond
// Date: 12 Jan, 2018
package word

import "testing"

// 功能测试函数，利用表进行构建
func TestIsPalindrome(t *testing.T) {
  var tests = []struct {
    input string
    want bool
  }{
    {"", true},
    {"a", true},
    {"aa", true},
    {"ab", false},
    {"kayak", true},
    {"detartrated", true},
    {"A man, a plan, a canal: Panama", true},
    {"Evil I did dwell; lewd did I Live.", true},
    {"Able was I ere I saw Elba", true},
    {"été", true},
    {"Et se resservir, ivresse reste.", true},
    {"palindrome", false},
    {"desserts", false},
  }
  for _, test := range tests {
    if got := IsPalindrome(test.input); got != test.want {
      t.Errorf("IsPalindrome(%q) = %v", test.input, got)
    }
  }
}
