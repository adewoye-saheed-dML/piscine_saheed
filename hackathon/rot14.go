package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, char := range runes {
		if char >= 'a' && char <= 'z' {
			runes[i] = (char-'a'+14)%26 + 'a'
		} else if char >= 'A' && char <= 'Z' {
			runes[i] = (char-'A'+14)%26 + 'A'
		}
	}
	return string(runes)
}
