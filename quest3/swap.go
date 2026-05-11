package piscine

func Swap(a *int, b *int) {
	y := *a
	*a = *b
	*b = y
}
