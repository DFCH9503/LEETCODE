package main

import(
	"fmt"
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64{
	concatSlices := slices.Concat(nums1, nums2)
	slices.Sort(concatSlices)
	lenSlice := len(concatSlices)
	midLen := lenSlice / 2

	if lenSlice % 2 == 0{
        return float64(concatSlices[midLen] + concatSlices[midLen - 1]) / 2
    }else{
        return float64(concatSlices[midLen])
    }

} 

func main(){
	nums1 := []int{}
	nums2 := []int{2, 3}
	fmt.Printf("The median of the two sorted arrays %v, %v is: %f", nums1, nums2, findMedianSortedArrays(nums1, nums2))
}

