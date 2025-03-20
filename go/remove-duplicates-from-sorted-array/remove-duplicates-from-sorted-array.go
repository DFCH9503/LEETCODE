package main

import(
	"fmt"
)

func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i <len(nums); i++{
		if nums[i] != nums[i-1]{
			nums [index] = nums[i]
			index++
		}
	}


	return index
}

func main(){
	array1 := []int {1 , 1 , 2, 2, 3, 4, 4}
	fmt.Println("answer to the first example is:", removeDuplicates(array1))
}