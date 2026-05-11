package piscine

// Helper function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	// If nb is less than or equal to 2, the next prime is always 2.
	// This covers negative numbers and 0, 1, 2.
	if nb <= 2 {
		return 2
	}

	// Start from nb and check every number upwards.
	for i := nb; ; i++ {
		if isPrime(i) {
			return i
		}

		// If we exceed the maximum possible integer, stop.
		if i < 0 {
			break
		}
	}
	return 0
}
