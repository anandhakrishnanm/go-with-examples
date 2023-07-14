package main

import (
	"os"
	"fmt"
	"bufio"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\nworld\n")
	err := os.WriteFile("hello.txt", d1, 0644)
	check(err)

	f, err := os.Create("hello1")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("write\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()
	
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("bufferd\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}