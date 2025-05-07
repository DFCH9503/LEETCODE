package main

import (
	"fmt"
	"slices"
)

func absInt(a int)int{
	if a < 0 {
		a *= -1
	}
	return a
}

func threeSumClosest(nums []int, target int) int{
	res := nums[0] + nums[1] + nums[2]
	slices.Sort(nums)

	for i := 0; i < len(nums); i++{
		j := i + 1
		k := len(nums) - 1
		for j < k{
			total := nums[i] + nums[j] + nums[k]
			if absInt(total - target) < absInt(res - target) {
				res = total
			}
			if total < target{
				j++
			}else if total > target{
				k--
			}else {
				return total
			}
		}
	}
	return res
}

func main(){
	nums := []int{-1,2,1,-4}
	target := 1

	fmt.Printf("for the array %v and the target %d, the closest sum to target is %d \n", nums, target, threeSumClosest(nums, target))
}