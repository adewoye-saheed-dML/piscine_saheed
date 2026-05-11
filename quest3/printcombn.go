package piscine

import "github.com/01-edu/z01"

var firstPrinted = true

func PrintCombN(n int) {
	firstPrinted = true
	if n <= 0 || n >= 10 {
		return
	}

	comb := make([]int, n)
	generateComb(0, 0, n, comb)
	z01.PrintRune('\n')
}

func generateComb(pos int, start int, n int, comb []int) {
	if pos == n {
		printComb(comb)
		return
	}

	for i := start; i <= 10-(n-pos); i++ {
		comb[pos] = i
		generateComb(pos+1, i+1, n, comb)
	}
}

func printComb(comb []int) {
	if !firstPrinted {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	firstPrinted = false

	for _, d := range comb {
		z01.PrintRune(int32('0' + d))
	}
}
