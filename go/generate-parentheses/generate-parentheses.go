package main

import(
	"fmt"
)

func generateParenthesis(n int)[]string{

	var solve func(open, close int, temp string)
    ans := []string{}

    solve = func(open, close int, temp string) {
        if len(temp) == 2*n {
            ans = append(ans, temp)
            return
        }

        if open > 0 {
            solve(open-1, close, temp+"(")
        }
        if close > open {
            solve(open, close-1, temp+")")
        }
    }

    solve(n, n, "")
    return ans
}

func main(){

	n := 3

	fmt.Printf("With %d pairs of parentheses all the well formed couples are: %v", n, generateParenthesis(n))
}