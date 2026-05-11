package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	items := ""
	for _, char := range str {
		if char == ' ' {
			summary[items]++
			items = ""
		} else {
			items += string(char)
		}
	}
	summary[items]++
	return summary
}
