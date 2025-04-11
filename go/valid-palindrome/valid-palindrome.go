package main

import(
	"fmt"
	"strings"
	"unicode"
)

func cleanString(input string) string {
	var builder strings.Builder

	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			builder.WriteRune(char)
		}
	}

	return strings.ToLower(builder.String())
}
func isPalindrome(s string) bool{
	sClean := cleanString(s)
	l := 0
	r := len(sClean)-1

	for l < r{
		if sClean[l] != sClean[r]{
			return false
		} 
		l++
		r--
	}
	return true
}


func main(){
	s := "A man, a plan, a canal: Panama"
	fmt.Printf("the string %s is palindrome: %v \n",s ,isPalindrome(s))
}