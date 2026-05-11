package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	f, err := os.Open(args[0])
	if err != nil {
		return
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
}

// dfdfdaa
