package main

import "fmt"

type (
	Entry struct {
		ID string
		Name string
	}
	Entries = map[uint32]Entry
)

func main() {
	data := Entries{}  // 等于 data := map[uint32]Entry{}
	data[32] = Entry{ID: "subond", Name: "bond"}
	fmt.Println(data)
}
