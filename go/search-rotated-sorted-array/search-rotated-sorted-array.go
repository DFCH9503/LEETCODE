package main

import(
	"fmt"
	// "slices"
	"math"
)

func search(nums []int, target int) int{

	//Binary Search to use it uncomment line 6
	l, r := 0 , len(nums) - 1

	for l <= r{
		m := int(math.Floor(float64((l + r) / 2)))

		if nums [m] == target {
			return m
		}else if nums [m] >= nums[l] {
			if nums[l] <= target && target <= nums[m] {
                r = m - 1
            } else {
                l = m + 1
            }
			} else {
				if nums[m] <= target && target <= nums[r] {
                l = m + 1
            } else {
                r = m - 1
            }
		}
	}
	return -1
}
	

	//Using in-built go function, to use it uncomment line 5
    // return slices.Index(nums, target)


func main(){
	nums :=  []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

fmt.Printf("The number = %d is in the array %v at the index %d, (zero index count), if answer is -1 the target isn't in the array", target, nums, search(nums, target))
}