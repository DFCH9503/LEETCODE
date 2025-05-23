function fourSum(nums, target){
    const res = []
    nums.sort((a, b) => a - b)
    const n = nums.length
    for (let i = 0; i < n - 3; i++) {
        if (i > 0 && nums[i] === nums[i - 1]) continue
        for (let j = i + 1; j < n - 2; j++) {
            if (j > i + 1 && nums[j] === nums[j - 1]) continue
            let l = j + 1, r = n - 1
            while (l < r) {
                const sum = nums[i] + nums[j] + nums[l] + nums[r]
                if (sum === target) {
                res.push([nums[i], nums[j], nums[l], nums[r]])
                while (l < r && nums[l] === nums[l + 1]) l++
                while (l < r && nums[r] === nums[r - 1]) r--
                l++
                r--
                } else if (sum < target) {
                l++
                } else {
                r--
                }
            }
        }
    }
    return res
}

let nums = [1,0,-1,0,-2,2], target = 0

console.log(`All the unique quadruplets of the array [${nums}] that added are equal to ${target} are `, fourSum(nums, target))