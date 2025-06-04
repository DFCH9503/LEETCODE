function findMedianSortedArrays(nums1, nums2) {
    let completeArray = nums1.concat(nums2).sort((a, b) => a - b)
    let middleComplete = completeArray.length / 2
    if(middleComplete % 2 == 0){
        return (completeArray[middleComplete] + completeArray[middleComplete - 1]) / 2
    }else{
        return completeArray[Math.floor(middleComplete)]
    }
}

let nums1 = [1, 2], nums2 = [3, 4]
console.log(`The median of the two sorted arrays ${nums1}, ${nums2} is: `, findMedianSortedArrays(nums1, nums2)) 