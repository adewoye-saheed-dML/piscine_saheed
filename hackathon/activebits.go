package piscine

func ActiveBits(n int) int {
	count := 0

	for n != 0 {

		count += n & 1

		n = int(uint(n) >> 1)
	}
	return count
}
