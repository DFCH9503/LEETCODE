//one way to do it but this consume to much memory because of the for...of cicle
// function searchInsert(nums, target){
//     let idx = 0
//     for (let value of nums){
//         if (value >= target){
//             return idx
//         }
//         idx++
//     }
//     return idx
// }

//other way to do it
function searchInsert(nums, target){
    let result = 0
    
    if(nums.includes(target)){
        result = nums.indexOf(target)
    }else{
        nums.push(target)
        nums.sort((a,b) => a - b)
        result = nums.indexOf(target)
    }
    return result
}


let nums = [1 , 3 , 5 , 6]
let target = 5
console.log("answer to the first example is:",searchInsert(nums, target))
