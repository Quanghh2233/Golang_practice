package main

import (
	"fmt"
	"os"
)

func main() {
	f := creatFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func creatFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Println(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
