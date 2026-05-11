package piscine

func LastRune(s string) rune {
	result := []rune(s)
	i := len(s) - 1
	return result[i]
}
