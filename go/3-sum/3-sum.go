package main

import(
	"fmt"
	"slices"
)

func threeSum(nums []int)[][]int{
	res := [][]int{}
	slices.Sort(nums)

	for i := 0; i < len(nums); i++{
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k{
			total := nums[i] + nums[j] + nums[k]
			if total > 0 {
				k--
			}else if total < 0{
				j++
			}else{
				newRow := []int{nums[i], nums[j], nums [k]}
				res = append(res, newRow)
				j++

				for nums[j] == nums[j-1] && j < k{
					j++
				}
			}
		}
	}
	return res
}

func main(){
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Printf("for the array %v, the possible ways to sum 0 with differents indexes are %v", nums, threeSum(nums))
}