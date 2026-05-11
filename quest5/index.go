package piscine

// import "fmt"

func Index(s string, toFind string) int {
	n := len(s)
	m := len(toFind)

	if m == 0 {
		return 0
	}

	for i := 0; i <= n-m; i++ {
		if s[i:i+m] == toFind {
			return i
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "alu"))
// }
