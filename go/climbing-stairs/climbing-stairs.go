package main

import(
	"fmt"

)

func climbStairs(n int) int {
	secondLast, last := 0, 1
    for i:=1; i<=n; i++ {
        secondLast, last = last, secondLast + last
    }
    return last
}

func main(){
	n := 2
	fmt.Println("answer to the first example is:",climbStairs(n))
}