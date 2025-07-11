package main

import(
	"fmt"
)

func removeElement(nums []int, val int) int{
	result := 0
	for i, n := range nums{
		if n != val{
			nums[result] = nums[i]
			result++
		}
	} 
	return result
} 

func main(){
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println("answer to the first example is:",removeElement(nums, val))
}