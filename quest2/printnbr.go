package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Special case for the smallest possible integer to avoid overflow
	if n == -9223372036854775808 || n == -2147483648 {
		if n < -3000000000 { // 64-bit MinInt
			z01.PrintRune('-')
			PrintNbr(9)
			PrintNbr(223372036854775808)
		} else { // 32-bit MinInt
			z01.PrintRune('-')
			PrintNbr(2)
			PrintNbr(147483648)
		}
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n > 9 {
		PrintNbr(n / 10)
	}

	z01.PrintRune(rune('0' + (n % 10)))
}
