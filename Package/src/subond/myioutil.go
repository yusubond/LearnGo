package subond

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func MyIO() {
	size, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes in file: %s", len(size), size)
}

func MyDir() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("./", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}

func MyFile() {
	content := []byte("HELLO SUBOND.")
	tmpfile, err := ioutil.TempFile("./", "tempfile")
	if err != nil {
		log.Fatal(err)
	}
	// defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
