package main

import(
	"fmt"
	"slices"
)

func fourSum(nums []int, target int)[][]int{
	res := [][]int{}
	n := len(nums)
	slices.Sort(nums)
	for i := 0; i < n-3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        for j := i + 1; j < n-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }

            l, r := j+1, n-1
            for l < r {
                currentSum := nums[i] + nums[j] + nums[l] + nums[r]

                if currentSum == target {
                    res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

                    l++
                    r--
                    for l < r && nums[l] == nums[l-1] {
                        l++
                    }
                    for l < r && nums[r] == nums[r+1] {
                        r--
                    }

                } else if currentSum < target {
                    l++
                } else {
                    r-- 
                }
            }
        }
    }

    return res
}

func main(){
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0

	fmt.Printf("All the unique quadruplets of the array %v that added are equal to %d are %v ", nums, target, fourSum(nums, target))
}