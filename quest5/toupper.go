package piscine

// import "fmt"

func ToUpper(s string) string {
	char := []rune(s)
	for i := 0; i < len(char); i++ {
		if char[i] >= 'a' && char[i] <= 'z' {
			char[i] -= 32
		}
	}
	return string(char)
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// }
