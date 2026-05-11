package piscine

// import "fmt"

func Concat(str1 string, str2 string) string {
	result := ""
	con := str1 + str2
	result += string(con)
	return result
}

// func main() {
// 	fmt.Println(Concat("Hello!", " How are you?"))
// }
