package piscine

// SortWordArr sorts a slice of strings in ascending ASCII order in place.
func SortWordArr(a []string) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Go compares strings lexicographically by their byte (ASCII) values
			if a[j] > a[j+1] {
				// Swap the elements
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
