package piscine

func Sqrt(nb int) int {
	for i := 1; i <= nb/2+1; i++ {
		square := i * i

		if square == nb {
			return i
		}

		if square > nb || square < 0 {
			return 0
		}
	}
	return 0
}
