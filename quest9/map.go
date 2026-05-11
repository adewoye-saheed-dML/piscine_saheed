package piscine

func Map(f func(int) bool, a []int) []bool {
	// Pre-allocate the result slice with the exact length needed for efficiency
	result := make([]bool, len(a))

	// Iterate through the input slice
	for i, value := range a {
		// Apply the function 'f' to the value and store it at the same index
		result[i] = f(value)
	}

	return result
}
