package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	max := a[0]
	for _, val := range a {
		if val > max {
			max = val
		}
	}

	return max
}
