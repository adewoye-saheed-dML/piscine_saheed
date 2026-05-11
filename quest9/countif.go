package piscine

func CountIf(f func(string) bool, tab []string) int {
	// Initialize a counter to keep track of the matches
	count := 0

	// Iterate through each string in the slice
	for _, str := range tab {
		// If the function 'f' evaluates to true for the current string, increment the counter
		if f(str) {
			count++
		}
	}

	// Return the final count
	return count
}
