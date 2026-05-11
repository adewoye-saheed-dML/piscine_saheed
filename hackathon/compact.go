package piscine

const N = 6

func Compact(ptr *[]string) int {
	temp := []string{}
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			temp = append(temp, (*ptr)[i])
		}
	}
	*ptr = []string{}
	for _, char := range temp {
		*ptr = append(*ptr, char)
	}
	return len(*ptr)
}
