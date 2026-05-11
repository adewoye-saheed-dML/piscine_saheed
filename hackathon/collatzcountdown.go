package piscine

// import "fmt"

func CollatzCountdown(start int) int {
	if start == 0 || start < 0 {
		return -1
	}

	count := 0

	for start != 1 {
		if start%2 == 0 {
			start = start / 2
			count++
		} else if start%2 != 0 {
			start = 3*start + 1
			count++
		}
	}

	return count
}

// func main() {
// 	steps := CollatzCountdown(12)
// 	fmt.Println(steps)
// }
