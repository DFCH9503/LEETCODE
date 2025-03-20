function twoSum(nums, target){
    let numToIndexMap={}
    for(let i=0 ; i<nums.length ; i++){
        let diff=target-nums[i]
        if (diff in numToIndexMap){
            return [i, numToIndexMap[diff]]
        }
        numToIndexMap[nums[i]]=i
    }
    return []
}

console.log("answer to the first example is:",twoSum([3, 2, 4], 6))
