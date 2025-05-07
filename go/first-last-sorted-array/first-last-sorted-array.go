package main

import (
	"fmt"
	"math"
)

func binarysearch(nums []int, target int, isSearchingLeft bool) int{
	left := 0
	right := len(nums) - 1
	idx := -1

	for left <= right{
		mid := int(math.Floor(float64((left + right) / 2)))
		
		if nums[mid] > target {
			right = mid - 1
		}else if nums[mid] < target{
			left = mid + 1
		}else{
			idx = mid
			if isSearchingLeft{
				right = mid - 1
			}else{
				left = mid +1
			}
		}
	}
	return idx
}

func searchRange(nums []int, target int)[]int{
	left := binarysearch(nums, target, true)
	right := binarysearch(nums, target, false)

	return []int{left, right}
}

func main(){
	nums  := []int{5,7,7,8,8,10}
	target := 8

	fmt.Printf("for the int %d the first and last indexes in slice %v are %v", target, nums, searchRange(nums, target))
}