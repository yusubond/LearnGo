// Name: max_test.go
// Author: subond
// Date: 13 Jan, 2018
package max

import "testing"

func TestMax(t *testing.T) {
  var tests = []struct{
    a, b, want int
  }{
    {1, 2, 2},
    {-8, 0, 0},
    {-1, -3, -1},
    {8, 19, 12},  // a error test case
  }

  for _, test := range tests {
    if got := Max(test.a, test.b); got != test.want {
      t.Errorf("Max(%d, %d) = %v", test.a, test.b, got)
    }
  }
}
