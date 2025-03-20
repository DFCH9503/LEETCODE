package main

import(
	"fmt"
)

func searchInsert(nums []int, target int) int{
	idx := 0

	for _, v := range nums {
		if v >= target {
			return idx
		}
		idx++
	}

	return idx

}

func main(){

	nums := []int{1 , 3 , 5 , 6}
	target := 5

	fmt.Println("answer to the first example is:",searchInsert(nums, target))
}