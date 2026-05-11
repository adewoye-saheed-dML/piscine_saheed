package piscine

// import "fmt"

func IsNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsNumeric("010203"))
// 	fmt.Println(IsNumeric("01,02,03"))
// }
