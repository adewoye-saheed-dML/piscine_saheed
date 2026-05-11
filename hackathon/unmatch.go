package piscine

// import "fmt"

func Unmatch(a []int) int {
	freq := make(map[int]int)

	for _, num := range a {
		freq[num] = freq[num] + 1
	}

	for _, num := range a {
		if freq[num]%2 != 0 {
			return num
		}
	}
	return -1
}

// func main() {
// 	a := []int{1, 2, 3, 1, 2, 3, 4,4,6}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
