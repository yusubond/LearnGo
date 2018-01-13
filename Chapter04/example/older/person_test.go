// Name: person.go
// Func:
// Author: subond
// Date: 13 Jan, 2018
package person

import "testing"

tests := []Person{{"Tom", 12}, {"Jack", 14}, {"John", 23}}

func TestOlder(t *testing.T) {
  for i := 1; i < 3; i++ {
      p, age := Older(tests[i-1], tests.[i])
      fmt.Printf("Of %s and %s, %s is older by %d years\n", tests[i-1].name, tests[i].name, p.name, age)
    }
  }
}
