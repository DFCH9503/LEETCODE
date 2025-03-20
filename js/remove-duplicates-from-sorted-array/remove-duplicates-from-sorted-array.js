function removeDuplicates(nums) {
    let index = 1
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] !== nums[i - 1]) {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}

let array1 = [ 1 , 1 , 2, 2, 3, 4, 4 ]
	console.log("answer to the first example is:", removeDuplicates(array1))
