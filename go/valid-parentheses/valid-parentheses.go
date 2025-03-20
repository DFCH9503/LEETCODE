package main

import(
	"fmt"
)

func isValid(s string) bool {
    stack := []rune{} // Stack for opening brackets
    hash := map[rune]rune{')': '(', ']': '[', '}': '{'}
    
    for _, char := range s {
        if match, found := hash[char]; found {
            if len(stack) > 0 && stack[len(stack)-1] == match {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        } else {
            stack = append(stack, char)
        }
    }
    return len(stack) == 0 
}

func main(){
	input := "()"
	fmt.Println("answer to the first example is:",isValid(input))
}