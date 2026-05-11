package piscine

// import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	result := []rune(s)
	if n <= 0 || n > len(result) {
		return 0
	}
	return result[n-1]
}

// func main() {
// 	z01.PrintRune(NRune("Salut!", 2))
// }
// hh
// Write function that returns the nth rune of a string. If not possible, it returns 0.
