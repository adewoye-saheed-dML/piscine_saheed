package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	// Iterate through the slice, comparing adjacent elements
	for i := 0; i < len(a)-1; i++ {
		// If the first is strictly greater than the second, it's not ascending
		if f(a[i], a[i+1]) > 0 {
			ascending = false
		}
		// If the first is strictly less than the second, it's not descending
		if f(a[i], a[i+1]) < 0 {
			descending = false
		}
	}

	// It is sorted if it remains either fully ascending or fully descending
	return ascending || descending
}
