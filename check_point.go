package main

// func CheckNumber(arg string) bool {
// 	for _, char := range arg {
// 		if char >= '0' && char <= '9' {
// 			return true
// 		}

// 	}
// 	return false
// }

// func main() {
// 	fmt.Println(CheckNumber("123"))
// 	fmt.Println(CheckNumber("H3llo"))
// 	fmt.Println(CheckNumber("Hello"))
// }

// func CountAlpha(s string) int {
// 	count := 0

// 	for _, char := range s {
// 		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
// 			count++
// 		}

// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(CountAlpha("Hello 123"))
// 	fmt.Println(CountAlpha("H e l l o"))
// 	fmt.Println(CountAlpha("12345"))
// }

// func CountChar(str string, c rune) int {
// 	result := 0
// 	if len(str) == 0 {
// 		return 0
// 	}

// 	for _, char := range str {
// 		if char == c {
// 			result++

// 		}
// 	}
// 	return result
// }

// func main() {
// 	fmt.Println(CountChar("Hello World", 'l'))
// 	fmt.Println(CountChar("5", '5'))
// 	fmt.Println(CountChar("ola!", '4'))
// }

// func PrintIfNot(str string) string {
// 	if len(str) < 3 || len(str) == 0 {
// 		return "G\n"
// 	}
// 	return "Invalid input\n"
// }

// func main() {
// 	fmt.Print(PrintIfNot("ab"))
// 	fmt.Print(PrintIfNot("abc"))
// 	fmt.Print(PrintIfNot(""))
// 	fmt.Print(PrintIfNot("145"))
// }

// func RectPerimeter(w, h int) int {
// 	peri := 2 * (w + h)

// 	if w < 0 || h < 0 {
// 		return -1
// 	}

// 	return peri
// }

// func main() {
// 	fmt.Println(RectPerimeter(10, 2))
// 	fmt.Println(RectPerimeter(434343, 13))
// 	fmt.Println(RectPerimeter(10, -2))
// }

// func RetainFirstHalf(str string) string {
// 	if len(str) == 0 {
// 		return ""
// 	} else if len(str) == 1 {
// 		return str
// 	}

// 	p := len(str) / 2
// 	return str[:p]
// }

// func main() {
// 	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
// 	fmt.Println(RetainFirstHalf("A"))
// 	fmt.Println(RetainFirstHalf(""))
// 	fmt.Println(RetainFirstHalf("Hello World"))
// }

// func DigitLen(n, base int) int {
// 	result := 0
// 	if base < 2 || base > 36 {
// 		return -1
// 	}

// 	if n < 0 {
// 		n = -n
// 	}

// 	for n != 0 {
// 		n = n / base
// 		result++
// 	}
// 	return result
// }

// func main() {
// 	fmt.Println(DigitLen(100, 10))
// 	fmt.Println(DigitLen(100, 2))
// 	fmt.Println(DigitLen(-100, 16))
// 	fmt.Println(DigitLen(100, -1))
// }

// func FirstWord(s string) string {
// 	result := ""
// 	i := 0
// 	for i < len(s) && s[i] == ' ' {
// 		i++
// 	}
// 	for _, char := range s[i:] {
// 		if char == ' ' {
// 			break
// 		}
// 		result += string(char)
// 	}
// 	return result + "\n"
// }

// func main() {
// 	fmt.Print(FirstWord("Hello World"))
// 	fmt.Print(FirstWord("I am a student"))
// 	fmt.Print(FirstWord("a"))
// 	fmt.Print(FirstWord("   lorem ipsum"))
// }

// func Gcd(a, b uint) uint {
// 	if a == 0 || b == 0 {
// 		return 0
// 	}

// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	return a
// }

// func main() {
// 	fmt.Println(Gcd(42, 10))
// 	fmt.Println(Gcd(42, 12))
// 	fmt.Println(Gcd(14, 77))
// 	fmt.Println(Gcd(17, 3))
// }

import (
	"fmt"
)

// func fpr(nb int) int {
// 	if nb < 2 {
// 		return 0
// 	}
// 	for current := nb; current >= 2; current-- {
// 		isPrime := true
// 		for i := 2; i*i <= current; i++ {
// 			if current%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			return current
// 		}
// 	}
// 	return 0
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println(0)
// 		return
// 	}
// 	n, err := strconv.Atoi(os.Args[1])
// 	if err != nil || n <= 0 {
// 		fmt.Println(0)
// 		return
// 	}
// 	sum := 0
// 	for p := fpr(n); p >= 2; p = fpr(p - 1) {
// 		sum += p
// 	}
// 	fmt.Println(sum)
// }

// func CanJump(nums []uint) bool {
// 	if len(nums) == 0 {
// 		return false
// 	}
// 	last := uint(len(nums) - 1)
// 	pos := uint(0)
// 	for pos < last {
// 		steps := nums[pos]
// 		if steps == 0 {
// 			return false
// 		}
// 		pos += steps
// 		if pos > last {
// 			return false
// 		}
// 	}
// 	return true
// }

// func Chunk(slice []int, size int) {
// 	if size == 0 {
// 		fmt.Println()
// 		return
// 	}
// 	result := [][]int{}
// 	for i := 0; i < len(slice); i += size {
// 		end := i + size
// 		if end > len(slice) {
// 			end = len(slice)
// 		}
// 		result = append(result, slice[i:end])
// 	}
// 	fmt.Println(result)
// }

// func ConcatAlternate(slice1, slice2 []int) []int {
// 	larger, smaller := slice1, slice2
// 	if len(slice2) > len(slice1) {
// 		larger, smaller = slice2, slice1
// 	}
// 	result := []int{}
// 	for i := 0; i < len(larger); i++ {
// 		result = append(result, larger[i])
// 		if i < len(smaller) {
// 			result = append(result, smaller[i])
// 		}
// 	}
// 	return result
// }

func ConcatSlice(slice1, slice2 []int) []int {
	return append(slice1, slice2...)
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
