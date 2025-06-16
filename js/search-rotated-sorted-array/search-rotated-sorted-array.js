function search(nums, target){

    //Binary Search
    // let l = 0
    // let r = nums.length - 1

    // while (l <= r) {
    //     let m = Math.floor((l + r) / 2)

    //     if (nums[m] === target) {
    //         return m
    //     } else if (nums[m] >= nums[l]) {
    //         if (nums[l] <= target && target <= nums[m]) {
    //             r = m - 1
    //         } else {
    //             l = m + 1
    //         }
    //     } else {
    //         if (nums[m] <= target && target <= nums[r]) {
    //             l = m + 1
    //         } else {
    //             r = m - 1
    //         }
    //     }
    // }

    // return -1 
    
    
    //Using in-built js function

    return nums.indexOf(target)
}

let nums = [4, 5, 6, 7, 0, 1, 2], target = 0

console.log(`The number = ${target} is in the array [${nums}] at the index`, search(nums, target), `(zero index count), if answer is -1 the target isn't in the array`)