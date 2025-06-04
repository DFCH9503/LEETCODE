function findMedianSortedArrays(nums1, nums2) {
    let completeArray = nums1.concat(nums2).sort((a, b) => a - b)
    let lenArray = completeArray.length, mid = lenArray / 2
    if(lenArray % 2 == 0){
        return (completeArray[mid] + completeArray[mid - 1]) / 2
    }else{
        return completeArray[Math.floor(mid)]
    }
}

let nums1 = [], nums2 = [2 ,3]
console.log(`The median of the two sorted arrays ${nums1}, ${nums2} is: `, findMedianSortedArrays(nums1, nums2)) 