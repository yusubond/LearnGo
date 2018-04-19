package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	readfile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 200)
	size, err := readfile.Read(data)

	writefile, err := os.Create("test_new.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = writefile.Write(data)
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("write %d bytes to file %s\n", size, writefile.Name())
}
