// Name: popcount.go
// Func: 包popcount统计数字中被置位的个数
// Author: subond
// Date: Dec 13, 2017
package popcount

var pc [256]byte

func init() {
  for i:= range pc {
    pc[i] = pc[i/2] + byte(i&1)
  }
}

// Popcount()返回数值x中被置为的个数
func Popcount(x uint64) int {
  return int(pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] +
    pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] + pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
}
