package main

import(
	"fmt"
	"slices"  //comment for the other method
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// last := m + n - 1
	// i := m - 1 //for nums1
	// j := n - 1 //for nums2
	
	// for j >=0 {
	// 	if i >= 0 && nums1[i] > nums2[j]{
	// 		nums1[last] = nums1[i]
	// 		i--
	// 	}else{
	// 		nums1[last]=nums2[j]
	// 		j--
	// 	}
	// 	last--
	// }
	
	// this solution needs extra memory, the slices package, and does not ude n parameter but it uses STL and can be checked with fmt.Println :)
	newNums1 := nums1[:m]
	newNums1 = append(newNums1, nums2...)
	slices.Sort(newNums1)
    return newNums1
}

func main(){
	nums1 := []int{1 , 2 , 3 , 0 , 0 , 0}
	m := 3 
	nums2 := []int{2 , 5 , 6}
	n := 3



	fmt.Println("answer to the first example is:",merge(nums1, m, nums2, n))
}