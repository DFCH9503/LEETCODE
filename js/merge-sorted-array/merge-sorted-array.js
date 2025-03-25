function merge(nums1, m, nums2, n){
    // let last = m + n - 1
	// let i = m - 1 //for nums1
	// let j = n - 1 //for nums2
	
	// for (j >=0){
	// 	if (i >= 0 && nums1[i] > nums2[j]){
	// 		nums1[last] = nums1[i]
	// 		i--
	// 	}else{
	// 		nums1[last]=nums2[j]
	// 		j--
	// 	}
	// 	last--
	// }


    nums1.length = m
    return nums1.concat(nums2).sort((a, b) => a-b)

}

let nums1 = [1 , 2 , 3 , 0 , 0 , 0], m = 3, nums2 =[2 , 5 , 6], n=3

console.log("answer to the first example is:",merge(nums1, m, nums2, n))