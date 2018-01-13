// Name: add_test.go
// Author: subond
// Date: 13 Jan, 2018
package add

import "testing"

var tests = []struct{
  a, want int
}{
  {1, 2},
  {12, 13},
  {-3, -2},
  {0, 1},
}

func TestAdd1(t *testing.T) {
  for _, test := range tests {
    if got := Add1(test.a); got != test.a {
      t.Errorf("Add1(%d) = %v", test.a, got)
    }
  }
}

func TestAdd2(t *testing.T) {
  for _, test := range tests {
    if got := Add2(&test.a); got != test.a {
      t.Errorf("Add2(%d) = %v", test.a, got)
    }
  }
}
