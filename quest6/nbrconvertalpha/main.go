package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) (int, bool) {
	res := 0
	if len(s) == 0 {
		return 0, false
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, false
		}
		res = res*10 + int(char-'0')
	}
	return res, true
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	isUpper := false
	startIndex := 0

	if args[0] == "--upper" {
		isUpper = true
		startIndex = 1
	}

	for i := startIndex; i < len(args); i++ {
		n, ok := BasicAtoi(args[i])

		if !ok || n < 1 || n > 26 {
			z01.PrintRune(' ')
		} else {
			if isUpper {
				z01.PrintRune(rune('A' + n - 1))
			} else {
				z01.PrintRune(rune('a' + n - 1))
			}
		}
	}

	if len(args) > startIndex {
		z01.PrintRune('\n')
	}
}
