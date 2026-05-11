package piscine

func AppendRange(min, max int) []int {
	// 1. Check if the range is invalid
	if min >= max {
		return nil
	}

	// 2. Declare a nil slice
	var result []int

	// 3. Loop from min to max (excluding max)
	for i := min; i < max; i++ {
		// 4. Dynamically grow the slice
		result = append(result, i)
	}

	return result
}
