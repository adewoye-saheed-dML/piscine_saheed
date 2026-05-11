package piscine

// import "fmt"

func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j = j - 1
		}
		nums[j+1] = key
	}
	return nums
}

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	insertionSort(nums)
	return nums[2]
}

// func main() {
// 	middle := Abort(2, 3, 8, 5, 7)
// 	fmt.Println(middle)
// }
