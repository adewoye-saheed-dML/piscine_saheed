package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	lastSlash := -1
	for i, char := range name {
		if char == '/' {
			lastSlash = i
		}
	}

	programName := name[lastSlash+1:]

	for _, char := range programName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
