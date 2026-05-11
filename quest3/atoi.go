package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	if len(s) == 0 {
		return 0
	}

	for _, char := range s {

		if char < '0' || char > '9' {
			return 0
		}

		digit := int(char - '0')
		result = result*10 + digit
	}

	return result * sign
}
