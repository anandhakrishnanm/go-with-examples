package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	check(err)
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile("name", d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent/child") 

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.Type().IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	fmt.Println("Listing subdir/parent/child")

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../../")
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
	check(err)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(" ", p, info.IsDir())
	return nil

}