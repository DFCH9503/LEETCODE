//using native js methods and functions

// function isPalindrome(x){

//     res = parseInt(x.toString().split("").reverse().join(""))

//     return res == x
// }

//using other approach, two pointer

function isPalindrome(x){
    let left = 0
    let right = x.toString().length
    let xString = x.toString()

    for (i = left; i < Math.floor(right/2) ; i++){
        let charRight = xString.charAt(i)
        let charLeft = xString.charAt(right-(i+1))
        if(charRight == charLeft){
            continue
        }else{
            return false
        }
    }
    
    return true
    
}

console.log("answer to the first example is:", isPalindrome(121))
