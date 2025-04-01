package main

import(
	"fmt"
	"math"
)

func longestPalindrome(s string)string{
	start := 0
	end := 0
	if  len(s) < 1 {
		return ""
	}
	for i := 0; i < len(s); i++{
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := math.Max(float64(len1), float64(len2))
		if int(maxLen) > end-start{
			start = i - (int(maxLen)-1) / 2
            end = i + int(maxLen)/2
		}
	}
	
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int{
	for left >= 0 && right < len(s) && s[left] == s[right]{
		left--
		right++
	}
	return right-left-1 
}

func main(){
	s:= "babad"
	fmt.Println("answer to the first example is:",longestPalindrome(s))
}