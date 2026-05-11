package piscine

func Any(f func(string) bool, a []string) bool {
	// Iterate through each string in the slice
	for _, str := range a {
		// If the function 'f' returns true for the current string,
		// we immediately return true since at least one element satisfies the condition
		if f(str) {
			return true
		}
	}

	// If the loop finishes and no element returned true, return false
	return false
}
