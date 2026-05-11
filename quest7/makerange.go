package piscine

func MakeRange(min, max int) []int {
	// 1. Check for invalid range
	if min >= max {
		return nil
	}

	// 2. Calculate the required size
	size := max - min

	// 3. Allocate the slice with the exact size needed
	// make([]type, length)
	result := make([]int, size)

	// 4. Fill the slice using indices
	for i := 0; i < size; i++ {
		result[i] = min + i
	}

	return result
}
