package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	var currentWord string

	for _, char := range s {
		// Check if the character is a whitespace separator
		if char == ' ' || char == '\t' || char == '\n' {
			// If we have accumulated a word, add it to our slice and reset
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			// Append the character to the current word
			currentWord += string(char)
		}
	}

	// Catch the last word if the string doesn't end with a whitespace
	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}
