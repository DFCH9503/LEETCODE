package main

import(
	"fmt"
)

func letterCombinations(digits string)[]string{
	if digits == "" {
		return []string{}
	}

	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var output []string

	var backtrack func(combination string, nextDigits string)
	backtrack = func(combination string, nextDigits string) {
		if nextDigits == "" {
			output = append(output, combination)
		} else {
			letters := phoneMap[nextDigits[0]-'2']
			for _, letter := range letters {
				backtrack(combination+string(letter), nextDigits[1:])
			}
		}
	}

	backtrack("", digits)
	return output
}


func main(){
	digits := "23"

	fmt.Printf("for the digits %s in a phone number, all the possible combinations are %v", digits, letterCombinations(digits))
}