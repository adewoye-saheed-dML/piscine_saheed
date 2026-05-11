package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Convert the input number (nbr) from baseFrom to a base-10 integer
	baseFromLen := len(baseFrom)
	var decimalValue int

	for i := 0; i < len(nbr); i++ {
		char := nbr[i]
		// Find the index (value) of the current character in baseFrom
		index := 0
		for j := 0; j < baseFromLen; j++ {
			if baseFrom[j] == char {
				index = j
				break
			}
		}
		// Multiply current total by base and add the new digit's value
		decimalValue = decimalValue*baseFromLen + index
	}

	// Step 2: Convert the base-10 integer to the target base (baseTo)
	baseToLen := len(baseTo)

	// Handle the edge case where the number is zero
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	var result []byte

	// Build the new string in reverse order using modulo
	for decimalValue > 0 {
		remainder := decimalValue % baseToLen
		result = append(result, baseTo[remainder])
		decimalValue /= baseToLen
	}

	// Reverse the result byte slice to get the correct order
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
