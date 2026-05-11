package piscine

func ConcatParams(args []string) string {
	var result string

	for i, arg := range args {
		result += arg
		// Add a newline only if it's not the last element
		if i < len(args)-1 {
			result += "\n"
		}
	}

	return result
}
