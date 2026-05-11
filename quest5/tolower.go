package piscine

// import "fmt"

func ToLower(s string) string {
	char := []rune(s)
	for i := 0; i < len(char); i++ {
		if char[i] >= 'A' && char[i] <= 'Z' {
			char[i] += 32
		}
	}
	return string(char)
}
