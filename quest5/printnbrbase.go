package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	baseLen := 0
	for range base {
		baseLen++
	}

	if baseLen < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i, char := range base {
		if char == '+' || char == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < baseLen; j++ {
			if char == rune(base[j]) {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	if nbr < 0 {
		z01.PrintRune('-')
	} else {
		nbr = -nbr
	}

	recursivePrint(nbr, base, baseLen)
}

func recursivePrint(nbr int, base string, baseLen int) {
	if nbr <= -baseLen {
		recursivePrint(nbr/baseLen, base, baseLen)
	}

	digitIndex := -(nbr % baseLen)
	z01.PrintRune(rune(base[digitIndex]))
}
