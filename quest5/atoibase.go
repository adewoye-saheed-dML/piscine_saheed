package piscine

func AtoiBase(s string, base string) int {
	baseLen := 0
	for range base {
		baseLen++
	}

	if baseLen < 2 {
		return 0
	}

	for i, char := range base {
		if char == '+' || char == '-' {
			return 0
		}
		for j := i + 1; j < baseLen; j++ {
			if char == rune(base[j]) {
				return 0
			}
		}
	}

	result := 0
	for _, charS := range s {
		// Find the value of the character based on its position in base
		val := -1
		for i, charB := range base {
			if charS == charB {
				val = i
				break
			}
		}

		if val == -1 {
			return 0
		}

		result = result*baseLen + val
	}

	return result
}
