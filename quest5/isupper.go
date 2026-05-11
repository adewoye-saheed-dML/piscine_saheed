package piscine

// import "fmt"

func IsUpper(s string) bool {
	for _, char := range s {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsUpper("HLLO"))
// 	fmt.Println(IsUpper("HELLO!"))
// }
