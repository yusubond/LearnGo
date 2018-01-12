// Name: word.go
// Func: 回文字符串
// Author: subond
// Date: 12 Jan, 2018
package word

import "unicode"

// IsPalindrome判断一个字符串是否为回文字符串
// 利用字符序列进行比较，而不是字节序列
func IsPalindrome(s string) bool {
  var letters []rune
  for _, r := range s {
    if unicode.IsLetter(r) {
      letters = append(letters, unicode.ToLower(r))
    }
  }
  for i := range letters {
    if letters[i] != letters[len(letters) - 1 - i] {
      return false
    }
  }
  return true
}
