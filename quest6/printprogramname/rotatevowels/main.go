package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	var fullText []rune
	for i, arg := range args {
		for _, char := range arg {
			fullText = append(fullText, char)
		}
		if i < len(args)-1 {
			fullText = append(fullText, ' ')
		}
	}

	var vowels []rune
	for _, char := range fullText {
		if isVowel(char) {
			vowels = append(vowels, char)
		}
	}

	vIdx := len(vowels) - 1
	for i := 0; i < len(fullText); i++ {
		if isVowel(fullText[i]) {
			fullText[i] = vowels[vIdx]
			vIdx--
		}
	}

	for _, char := range fullText {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
