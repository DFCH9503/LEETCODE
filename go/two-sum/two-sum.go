package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    numToIndexMap := make(map[int]int)

    for i, num := range nums {
        
        diff := target - num

        if idx, found := numToIndexMap[diff]; found {
            return []int{i, idx}
        }

        numToIndexMap[num] = i
    }
	return nil
}

func main(){
	twoSum([]int{3, 2, 4}, 6)
	fmt.Println("answer to the first example is:", twoSum([]int{3, 2, 4}, 6))
}