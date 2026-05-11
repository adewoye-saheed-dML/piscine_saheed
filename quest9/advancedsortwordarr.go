package piscine

// AdvancedSortWordArr sorts a slice of strings based on a custom comparison function.
// The function f should return:
//   - a positive integer if a > b
//   - zero if a == b
//   - a negative integer if a < b
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// If f returns a value > 0, it means a[j] is "greater" than a[j+1]
			// according to the custom logic, so we swap them.
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
