package main

import(
	"fmt"

)

func mySqrt(x int) int {
	root := x
	for root*root > x{
		root = (root + x / root) / 2 
	}
	return root
}

func main(){
	x := 8
	fmt.Println("answer to the first example is:",mySqrt(x))
}