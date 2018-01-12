// Name: word_test.go
// Func: 随机测试
// Author: subond
// Date: 12 Jan, 2018
package word

import "testing"
import "math/rand"
import "time"

func randomPalindrome(rng *rand.Rand) string {
  n := rng.Intn(25)       // 随机字符串的最大长度为24
  runes := make([]rune, n)
  for i := 0; i < (n + 1)/2; i++ {
    r := rune(rng.Intn(0x1000))     // 随机字符最大是 '\u0999'
    runes[i] = r
    runes[n - 1 - i] = r
  }
  return string(runes)
}

// 功能测试函数，利用表进行构建
func TestIsPalindrome(t *testing.T) {
  // 初始化一个伪随机数生成器
  seed := time.Now().UTC().UnixNano()
  t.Logf("Random seed: %d", seed)
  rng := rand.New(rand.NewSource(seed))
  for i := 0; i < 1000; i++ {
    p := randomPalindrome(rng)
    if !IsPalindrome(p) {
      t.Errorf("IsPalindrome(%q) = false", p)
    }
  }
}
