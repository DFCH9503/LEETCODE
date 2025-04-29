function threeSumClosest(nums, target){
    let minDiff = Infinity, res = 0
    nums.sort((a, b) => a - b)
    for(let i = 0; i < nums.length; i++){
        let j = i + 1
        let k = nums.length - 1

        while(j < k){
            let total = nums[i] + nums[j] + nums [k]
            let currDiff = Math.abs(total - target)
            if(currDiff < minDiff){
                minDiff = currDiff
                res = total
            }
            if(total < target){
                j++ 
            }else{
                k--
            }
        }
    }
    return res
}

let nums = [-1,2,1,-4], target = 1
console.log("answer to the first example is:",threeSumClosest(nums, target))