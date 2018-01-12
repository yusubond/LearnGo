// Name: word_test.go
// Func: package word的测试文件
// Author: subond
// Date: 12 Jan, 2018
package word

import "testing"

// 功能测试函数
func TestPalindrome(t *testing.T) {
  if !IsPalindrome("detartrated") {
    t.Error(`IsPalindrome("detartrated") = false`)
  }
  if !IsPalindrome("kayak") {
    t.Error(`IsPalindrome("kayak") == false`)
  }
}

func TestNonPalindrome(t *testing.T) {
  if IsPalindrome("palindrome") {
    t.Error(`IsPalindrome("palindrome") = true`)
  }
}

func TestFrenchPalindrome(t *testing.T) {
  if !IsPalindrome("été") {
    t.Error(`IsPalindrome("été") = false`)
  }
}

func TestCanalPalindrome(t *testing.T) {
  input := "A man, a plan, a canal: Panama"
  if !IsPalindrome(input) {
    t.Errorf(`IsPalindrome(%q) = false`, input)
  }
}
