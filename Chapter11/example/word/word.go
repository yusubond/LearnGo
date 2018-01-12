// Name: word.go
// Func: 回文字符串
// Author: subond
// Date: 12 Jan, 2018
package word

// IsPalindrome判断一个字符串是否为回文字符串
func IsPalindrome(s string) bool {
  for i := range s {
    if s[i] != s[len(s) - i - 1] {
      return false
    }
  }
  return true
}
