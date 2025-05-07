function searchRange(nums, target) {
    const binarySearch = (nums, target, isSearchingLeft) => {
        let left = 0
        let right = nums.length - 1
        let idx = -1

        while (left <= right) {
            const mid = Math.floor((left + right) / 2)
            
            if (nums[mid] > target) {
                right = mid - 1
            } else if (nums[mid] < target) {
                left = mid + 1
            } else {
                idx = mid
                if (isSearchingLeft) {
                    right = mid - 1
                } else {
                    left = mid + 1
                }
            }
        }

        return idx
    }
    
    const left = binarySearch(nums, target, true)
    const right = binarySearch(nums, target, false)
    
    return [left, right] 
}

let nums = [5,7,7,8,8,10], target = 8

console.log("answer to the first example is:",searchRange(nums, target))