package main

import(
	"fmt"
	str "github.com/mgutz/str" //this answer uses this third part package, you must init the mod and require this package from github
)

func srtStr(haystack, needle string) int{
	return str.IndexOf(haystack , needle , 0)
}

func main(){

	haystack := "sadbutsad"
	needle := "sad"
	fmt.Println("answer to the first example is:",srtStr(haystack, needle))
}