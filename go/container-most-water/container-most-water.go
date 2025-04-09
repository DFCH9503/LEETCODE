package main

import(
	"fmt"
)

func getMin(a, b int)int{
	if a < b{
		return a
	}else{
		return b
	}
}

func getMax(a, b int)int{
	if a < b{
		return b
	}else{
		return a
	}
}

func maxArea(height []int) int {
    l := 0
	r := len(height)-1
	res := 0

	for l < r{
		currWidth := r-l
		currHeight := getMin(height[l], height[r])
		res = getMax(res, currWidth*currHeight)
		if height[l] <= height[r] {
            l++
        } else {
            r--
        }
	} 
	return res
}

func main(){
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println("answer to the first example is:",maxArea(height))
}