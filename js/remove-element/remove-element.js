function removeElement(nums, val){
    let result = 0
    for (let i = 0 ; i < nums.length ; i++){
        if (nums[i] != val){
            nums[result] = nums [i]
            result++
        }
    }
    return result
}

let nums = [3,2,2,3]
let val = 3
console.log("answer to the first example is:",removeElement(nums, val))