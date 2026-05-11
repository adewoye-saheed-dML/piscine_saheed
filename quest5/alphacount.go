package piscine

// import "fmt"

func AlphaCount(s string) int {
	result := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			result++
		}
	}
	return result
}

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := AlphaCount(s)
// 	fmt.Println(nb)
// }
