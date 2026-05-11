package piscine

// import "fmt"

func IsAlpha(s string) bool {
	for _, char := range s {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsAlpha("Hello! How are you?"))
// 	fmt.Println(IsAlpha("HelloHowareyou"))
// 	fmt.Println(IsAlpha("What's this 4?"))
// 	fmt.Println(IsAlpha("Whatsthis4"))
// }
