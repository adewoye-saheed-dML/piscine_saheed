package main

import (
	"os"
)

// ParseInt manually parses strings to int64 and detects value overflows.
func ParseInt(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	neg := false
	if s[0] == '-' {
		neg = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	if len(s) == 0 {
		return 0, false
	}

	var res uint64
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false // Invalid character
		}
		digit := uint64(s[i] - '0')

		// Check for 64-bit bounds overflow
		if !neg && res > (9223372036854775807-digit)/10 {
			return 0, false
		}
		if neg && res > (9223372036854775808-digit)/10 {
			return 0, false
		}
		res = res*10 + digit
	}

	if neg {
		return -int64(res), true
	}
	return int64(res), true
}

// PrintInt handles outputting the resulting int64 to standard output.
func PrintInt(n int64) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}
	var b [32]byte
	i := 31
	b[i] = '\n'
	neg := false

	// Convert using uint64 to safely handle math.MinInt64
	u := uint64(n)
	if n < 0 {
		neg = true
		u = -u
	}

	for u > 0 {
		i--
		b[i] = byte(u%10) + '0'
		u /= 10
	}

	if neg {
		i--
		b[i] = '-'
	}
	os.Stdout.Write(b[i:])
}

// Arithmetic functions with overflow detection
func Add(a, b int64) (int64, bool) {
	c := a + b
	if (c > a) != (b > 0) && b != 0 {
		return 0, false // Overflow
	}
	return c, true
}

func Sub(a, b int64) (int64, bool) {
	c := a - b
	if (c < a) != (b > 0) && b != 0 {
		return 0, false // Overflow
	}
	return c, true
}

func Mul(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	c := a * b
	if c/b != a {
		return 0, false // Overflow
	}
	// Edge case: math.MinInt64 * -1 causes an overflow panic in Go
	if (a == -9223372036854775808 && b == -1) || (b == -9223372036854775808 && a == -1) {
		return 0, false
	}
	return c, true
}

func Div(a, b int64) (int64, bool) {
	// Edge case: math.MinInt64 / -1 causes a panic in Go
	if a == -9223372036854775808 && b == -1 {
		return 0, false
	}
	return a / b, true
}

func Mod(a, b int64) (int64, bool) {
	if a == -9223372036854775808 && b == -1 {
		return 0, false
	}
	return a % b, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := ParseInt(os.Args[1])
	b, ok2 := ParseInt(os.Args[3])
	op := os.Args[2]

	if !ok1 || !ok2 {
		return
	}

	var res int64
	var ok bool

	switch op {
	case "+":
		res, ok = Add(a, b)
	case "-":
		res, ok = Sub(a, b)
	case "*":
		res, ok = Mul(a, b)
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		res, ok = Div(a, b)
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		res, ok = Mod(a, b)
	default:
		return
	}

	if !ok {
		return
	}

	PrintInt(res)
}
