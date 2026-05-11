package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	help := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n"
	for _, r := range help {
		z01.PrintRune(r)
	}
}

func sortString(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
	return string(r)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	var insertStr string
	var baseStr string
	isOrder := false

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			isOrder = true
		} else {
			baseStr = arg
		}
	}

	// Step 1: Combine strings
	finalStr := baseStr + insertStr

	// Step 2: Sort if requested
	if isOrder {
		finalStr = sortString(finalStr)
	}

	// Step 3: Print result
	for _, r := range finalStr {
		z01.PrintRune(r)
	}
	if len(finalStr) > 0 {
		z01.PrintRune('\n')
	}
}
