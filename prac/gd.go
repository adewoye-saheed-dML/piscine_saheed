// package main

// import (
// 	"fmt"
// )

// func digi(x, y int) int {
// 	result := 0
// 	if y <= 1 || y > 36 {
// 		return -1
// 	}
// 	if x < 0 {
// 		x = -1 * x
// 	}
// }

// // // // 	for x != 0 {
// // // // 		x= x / y
// // // // 		result+=1
// // // // 	}
// // // // 	return result
// // // // }

// // // // func main() {
// // // // 	fmt.Println(digi(100, -1))
// // // // }

// // // // func first(s string) string {
// // // // 	result:= ""
// // // // 	for _, char := range s {
// // // // 		if  char==' '{
// // // // 			break
// // // // 		} else {
// // // // 			result+=string(char)
// // // // 		}
// // // // 	}
// // // // 	return result + "\n"
// // // // }

// // // // func main() {
// // // //     fmt.Print(first("Hello World"))
// // // // }

// // // // func fish(x int) string {
// // // // 	if x<0 {
// // // // 		return "negative"
// // // // 	} else if (x%2==0) && (x%3==0) {
// // // // 		return "fish and chips"
// // // // 	} else if x%2==0 {
// // // // 		return "fish"
// // // // 	}else if  x%3==0 {
// // // // 		return "chips"
// // // // 	}
// // // // 	return "non divisible"
// // // // }

// // // // func main() {
// // // //     fmt.Println(fish(-2))
// // // // }

// // // // func gcd(a,b uint) uint {
// // // // 	if a==0 || b==0{
// // // // 		return 0
// // // // 	}
// // // // 	for b!=0{
// // // // 		a, b = b, a%b
// // // // 	}
// // // // 	return a
// // // // }

// // // // func main() { fmt.Println(findprime(10))
// // // //     fmt.Println(gcd(42, 10))
// // // //     fmt.Println(gcd(42, 12))
// // // //     fmt.Println(gcd(14, 77))
// // // //     fmt.Println(gcd(17, 3))
// // // // }

// // // // func hash(s string) string{
// // // // 	result:= ""
// // // // 	size := int32(len(s))
// // // // 	for _, char := range s {
// // // // 		has:= (int32(char) + size)%127
// // // // 		if has <32 {
// // // // 			has+=33
// // // // 		}
// // // // 		result+= string(has)
// // // // 	}
// // // // 	return result
// // // // }

// // // // func main() {
// // // //     fmt.Println(hash("A"))
// // // //     fmt.Println(hash("AB"))
// // // // 	fmt.Println(hash("Hello"))
// // // // }

// // // // func last(s string) string {
// // // // 	result := ""
// // // // 	i := len(s) - 1

// // // // 	for i >= 0 && s[i] == ' ' {
// // // // 		i--
// // // // 	}
// // // // 	for i >= 0 && s[i] != ' ' {
// // // // 		result = string(s[i]) + result
// // // // 		i--
// // // // 	}
// // // // 	return result + "\n"
// // // // }

// // // // func main() {
// // // // 	fmt.Print(last("Hello World"))
// // // // }

// // // // func RepeatAlpha(s string) string {
// // // // 	result := ""

// // // // 	for _, char := range s {
// // // // 		count := 1

// // // // 		if char >= 'a' && char <= 'z' {
// // // // 			count = int(char - 'a' + 1)
// // // // 		} else if char >= 'A' && char <= 'Z' {
// // // // 			count = int(char - 'A' + 1)
// // // // 		}

// // // // 		for i := 0; i < count; i++ {
// // // // 			result += string(char)
// // // // 		}
// // // // 	}

// // // // 	return result
// // // // }

// // // // func cts(s string) string{
// // // // 	var res []byte

// // // // 	for i, c := range []byte(s){
// // // // 		if (c<'a' || c>'z') && (c<'A' || c>'Z'){
// // // // 			return s
// // // // 		}
// // // // 		if c>='A' && c<='Z'{
// // // // 			if i==len(s)-1 || s[i+1]>='A' && s[i+1]<='Z'{
// // // // 				return s
// // // // 			}
// // // // 			if i>0{
// // // // 				res = append(res, '_')
// // // // 			}
// // // // 			c|=32
// // // // 		}
// // // // 		res=append(res,c)
// // // // 	}
// // // // 	return string(res)
// // // // }

// // // // fromto.go
// // // // package piscine

// // // // import "fmt"

// // // // func FromTo(from int, to int) string {
// // // // 	if from < 0 || from > 99 || to < 0 || to > 99 {
// // // // 		return "Invalid\n"
// // // // 	}
// // // // 	result := fmt.Sprintf("%02d", from)
// // // // 	step := 1
// // // // 	if from > to {
// // // // 		step = -1
// // // // 	}
// // // // 	for i := from + step; i != to+step; i += step {
// // // // 		result += fmt.Sprintf(", %02d", i)
// // // // 	}
// // // // 	return result + "\n"
// // // // }

// // // func fto(from, to int) string {
// // // 	if from < 0 || from > 99 || to < 0 || to > 99 {
// // // 		return "invalid\n"
// // // 	}

// // // 	result := fmt.Sprintf("%02d", from)

// // // 	step := 1
// // // 	if from > to {
// // // 		step = -1
// // // 	}

// // // 	for i := from + step; i != to+step; i += step {
// // // 		result += fmt.Sprintf(", %02d", i)
// // // 	}
// // // 	return result + "\n"
// // // }

// // // func main() {
// // // 	fmt.Print(fto(1, 10))
// // // 	fmt.Print(fto(10, 1))
// // // 	fmt.Print(fto(10, 10))
// // // 	fmt.Print(fto(100, 10))
// // // }

// // // func findprime(nb int) int {
// // // // 	if nb < 2 {
// // // // 		return 0
// // // // 	}

// // // // 	for current:=nb; current>=2; current--{
// // // // 		isPrime := true

// // // // 		for i:=2; i*i<= current; i++{
// // // // 			if current%i==0 {
// // // // 				isPrime=false
// // // // 				break
// // // // 			}
// // // // 		}
// // // // 		if isPrime{
// // // // 			return current
// // // // 		}
// // // // 	}
// // // // 	return 0
// // // // }

// // // func fpr(nb int)int{
// // // 	if nb < 2 {
// // // 		return 0
// // // 	}

// // // 	for current:=nb; current>=2; current-- {
// // // 		isprime:= true

// // // 		for i:=2; i*i<=current; i++{
// // // 			if current%i==0{
// // // 				isprime = false
// // // 				break
// // // 			}
// // // 		}
// // // 		if isprime{
// // // 			return current
// // // 			}

// // // 		}
// // // 		return 0
// // // 	}

// // // func main() {
// // //     fmt.Println(fpr(5))
// // //     fmt.Println(fpr(4))
// // // }

// // // // func IsCapitalized(s string) bool {
// // // // 	if len(s) == 0 {
// // // // 		return false
// // // // 	}
// // // // 	inWord := false
// // // // 	for i := 0; i < len(s); i++ {
// // // // 		if s[i] == ' ' {
// // // // 			inWord = false
// // // // 		} else if !inWord {
// // // // 			if s[i] >= 'a' && s[i] <= 'z' {
// // // // 				return false
// // // // 			}
// // // // 			inWord = true
// // // // 		}
// // // // 	}
// // // // 	return true
// // // // }

// // // func main() {
// // //     fmt.Println(cap("Hello! How are you?"))
// // //     fmt.Println(cap("Hello How Are You"))
// // //     fmt.Println(cap("Whats 4this 100K?"))
// // //     fmt.Println(cap("Whatsthis4"))
// // // fmt.Println(IsCapitalized("!!!!Whatsthis4"))
// // // fmt.Println(IsCapitalized(""))

// // // func Itoa(n int) string{

// // // 	result:= ""

// // // 	if n == 0{
// // // 		return "0"
// // // 	}

// // // 	isneg:= false
// // // 	if n<0{
// // // 		isneg = true
// // // 		n= -n
// // // 	}
// // // 	for n>0 {
// // // 		dig:= n%10
// // // 		ch:= string(dig+'0')
// // // 		result= ch+ result
// // // 		n = n/10
// // // 	}
// // // 	if isneg{
// // // 		result = "-" + result
// // // 	}
// // // 	return result
// // // }
// // // func main() {
// // //     fmt.Println(Itoa(12345))
// // //     fmt.Println(Itoa(0))
// // //     fmt.Println(Itoa(-1234))
// // //     fmt.Println(Itoa(987654321))
// // // }

// // // func main (){
// // // 	for i:='9'; i>='2'; i-- {
// // // 		for j:=i-1; j>='1'; j--{
// // // 			for k:=j-1; k>='0'; k--{

// // // 				z01.PrintRune(i)
// // // 				z01.PrintRune(j)
// // // 				z01.PrintRune(k)

// // // 			if i!='2' || j!='1' || k!='0'{
// // // 				z01.PrintRune(',')
// // // 				z01.PrintRune(' ')
// // // 			}
// // // 			}
// // // 		}
// // // 	}
// // // 	z01.PrintRune('\n')
// // // }

// // package main

// // import (
// // 	"fmt"
// // )

// // // func ThirdTimeIsACharm(str string) string {
// // // 	// If the string is empty or has fewer than 3 characters, return a newline
// // // 	if len(str) < 3 {
// // // 		return "\n"
// // // 	}

// // // 	result := ""

// // // 	// Start at index 2 (the 3rd character) and step by 3 each iteration
// // // 	for i := 2; i < len(str); i += 3 {
// // // 		result += string(str[i])
// // // 	}

// // // 	// Add the final newline as requested
// // // 	return result + "\n"
// // // }

// // // func main() {
// // // 	fmt.Print(ThirdTimeIsACharm("123456789"))
// // // 	fmt.Print(ThirdTimeIsACharm(""))
// // // 	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
// // // 	fmt.Print(ThirdTimeIsACharm("12"))
// // // }

// // func unique(str1, str2 string) int {
// // 	if str1 == "" && str2 == "" {
// // 		return -1
// // 	}

// // 	set1 := make(map[rune]bool)
// // 	set2 := make(map[rune]bool)

// // 	for _, ch := range str1 {
// // 		set1[ch] = true
// // 	}
// // 	for _, ch := range str2 {
// // 		set2[ch] = true
// // 	}

// // 	count := 0
// // 	for ch := range set1 {
// // 		if !set2[ch] {
// // 			count++
// // 		}
// // 	}

// // 	for ch := range set2 {
// // 		if !set1[ch] {
// // 			count++
// // 		}
// // 	}
// // 	return count
// // }

// // func main() {
// // 	fmt.Println(unique("foo", "boo")) // 2
// // 	fmt.Println(unique("", ""))       // -1
// // 	fmt.Println(unique("abc", "def")) // 6
// // }

// // func zip(s string) string {
// // 	if s == "" {
// // 		return "s"
// // 	}

// // 	result := ""
// // 	count := 1

// // 	for i := 1; i < len(s); i++ {
// // 		if s[i] == s[i-1] {
// // 			count++
// // 		} else {
// // 			result += fmt.Sprintf("%d%c", count, s[i-1])
// // 			count = 1
// // 		}
// // 	}
// // 	result += fmt.Sprintf("%d%c", count, s[len(s)-1])
// // 	return result + "s"
// // }

// func fpr( nb int) int {
// 	if nb <2 {
// 		return 0
// 	}

// 	for current:=0; current>=2; current--{
// 		isprime:= true

// 		for i:=2; i*i<=current; i++{
// 			if current%i==0{
// 				isprime= false
// 				break
// 			}

// 			if isprime{
// 				return current
// 			}
// 		}
// 	}
// 	return 0
// }

// func fto(from, to int)string{
// 	if from<0 || from>99 || to<0 || to>99{
// 		return "invalid\n"
// 	}
// 	result:= fmt.Sprintf{"%02d", from}

// 	step:=1
// 	if from>to{
// 		step=-1
// 	}

// 	for i:=from+step; i<to+step; i+=step{
// 		result= fmt.Sprintf(", %02d", i)
// 	}
// 	return result
// }

// func iscap(s string) bool{
// 	if len(s)==0{
// 		return false
// 	}

// 	inword:=false

// 	for i:=0; i<len(s); i++{
// 		if s[i]==' '{
// 			return false
// 		}else if !inword{
// 			if s[i]>='a' && s[i]<='z'{
// 			return false
// 		}
// 		inword= true
// 	}

// 	}
// 	return true
// }

// func itoa(n int)string{
// 	if n ==0{
// 		return "0"
// 	}

// 	isneg:=false
// 	result:= ""

// 	if n<0{
// 		isneg=true
// 		n=-n
// 	}
// 	if n>0{
// 		dig:= n%10
// 		ch:= string(dig+'0')
// 		result+=ch
// 		n=n/10
// 	}
// 	if isneg{
// 		result="-" + result
// 	}
// 	return result
// }

// fun main(){
// 	for i:='9' ; i>='2'; i++{
// 		for j:=i-1; j>='1'; j++{
// 			for k:=j-1; k>='0'{

// 				z01.PrintRune(i)
// 				z01.PrintRune(j)
// 				z01.PrintRune(k)

// 			if i!='2' || j!='1'|| k!='0'{
// 				z01.PrintRune(',')
// 				z01.PrintRune(' ')
// 			}

// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func third(str string)string{
// 	if len(str)<3{
// 		return "\n"
// 	}

// 	result:=""

// 	for i:=2; i<len(str); i+=3{
// 		result= string(str[i]) +result
// 	}
// 	return result + "\n"
// }

// func unique(str1, str2) int{
// 	if str1=="" && str2==""{
// 		return -1
// 	}

// 	set1:= make(map[rune]bool)
// 	set2:=make(map[rune]bool)

// 	for _, ch := range str1{
// 		set1[ch]=true
// 	}

// 	for _, ch := range str2{
// 		set2[ch]=true
// 	}

// 	count:=0
// 	f0r ch := range set1{
// 		if !set2[ch]{
// 			count++
// 		}
// 	}
// 	f0r ch := range set2{
// 		if !set1[ch]{
// 			count++
// 		}
// 	}
// 	return count
// }

// func zip(s string) string{
// 	if s == ""{
// 		return "$"
// 	}

// 	result:=""
// 	count:=1

// 	for i:=1; i<len(s); i++{
// 		if s[i]==s[i-1]{
// 			count++
// 		} else{
// 			result+=fmt.Sprintf("%d%c", count, s[i-1])
// 			count=1
// 		}
// 	}
// 	result+=fmt.Sprintf("%d%c", count, s[len(s)-1])
// 	return result
// }

// Empty file

package main

import (
	"fmt"
	"strconv"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := strconv.Itoa(from)

	step := -1
	if from < to {
		step = 1
		result = "0" + strconv.Itoa(from)
	}

	for i := from + step; i != to+step; i += step {
		conv := strconv.Itoa(i)

		if i < 10 {
			conv = ", 0" + strconv.Itoa(i)
		} else if i == to {
			conv = ", " + strconv.Itoa(i)
		}
		result += conv
	}

	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
