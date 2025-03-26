package main

import(
	"fmt"
)

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func lengthLongestSubstring(s string) int {
    mp := make(map[byte]int)
	maxi := 0
	left := 0
	for right:=0; right<len(s); right++ {
		if val, exists := mp[s[right]]; exists {
			left = max(left, val+1)
	}
		mp[s[right]] = right
		maxi = max(maxi, right-left+1)
	}
	return maxi
}

func main(){
	s := "abcabcbb"
	fmt.Println("answer to the first example is:",lengthLongestSubstring(s))
}