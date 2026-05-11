package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isStartOfWord := true

	for i := 0; i < len(runes); i++ {
		if isAlphanumeric(runes[i]) {
			if isStartOfWord {
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] -= 32
				}
				isStartOfWord = false
			} else {
				// Lowercase if it's an uppercase letter
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] += 32
				}
			}
		} else {
			isStartOfWord = true
		}
	}
	return string(runes)
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
