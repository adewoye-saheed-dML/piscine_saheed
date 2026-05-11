package piscine

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)

	// Safeguard for an empty separator
	if sepLen == 0 {
		for _, char := range s {
			result = append(result, string(char))
		}
		return result
	}

	start := 0
	// Loop through the string 's'
	for i := 0; i <= len(s)-sepLen; i++ {
		// Check if the substring matches the separator
		if s[i:i+sepLen] == sep {
			// Append the word found before the separator
			result = append(result, s[start:i])
			// Update the starting point for the next word
			start = i + sepLen
			// Jump the index forward to skip the rest of the separator
			i = i + sepLen - 1
		}
	}

	// Append the remaining part of the string after the last separator
	result = append(result, s[start:])

	return result
}
