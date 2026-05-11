package piscine

func ReverseMenuIndex(menu []string) []string {
	n := len(menu)

	result := make([]string, n)

	for i := 0; i < n; i++ {
		result[n-1-i] = menu[i]
	}

	return result
}
