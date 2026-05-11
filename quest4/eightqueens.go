package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	// positions[col] = row_index
	var positions [8]int
	solve(0, &positions)
}

func solve(col int, positions *[8]int) {
	// Base Case: If we reached column 8, we found a valid solution
	if col == 8 {
		printSolution(positions)
		return
	}

	// Try placing a queen in each row (1 to 8) for the current column
	for row := 1; row <= 8; row++ {
		if isSafe(col, row, positions) {
			positions[col] = row
			solve(col+1, positions) // Move to the next column
		}
	}
}

func isSafe(col, row int, positions *[8]int) bool {
	for i := 0; i < col; i++ {
		prevRow := positions[i]

		// 1. Check if same row
		if prevRow == row {
			return false
		}

		// 2. Check diagonals
		// If the horizontal distance equals the vertical distance, they are diagonal
		diff := col - i
		if prevRow == row-diff || prevRow == row+diff {
			return false
		}
	}
	return true
}

func printSolution(positions *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(positions[i] + '0'))
	}
	z01.PrintRune('\n')
}
