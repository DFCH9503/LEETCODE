package main

import(
	"fmt"
)

func lengthOfLastWord(s string) int{
	res := 0
	for i:= len(s)-1 ; i >= 0 ; i--{
		if s[i] != 32{  //32 is ascii for space
			res++
		}
		if res >=1 && s[i]==32{
			return res
		}
	}
	return res
}

func main(){

	s := "luffy is still joyboy"
	fmt.Println("answer to the first example is:",lengthOfLastWord(s))
}