function cleanString(input) {
    return input.replace(/[^a-zA-Z0-9]/g, '');
}

function validPalindrome(s){
    let sClean = cleanString(s).toLowerCase().split("")
    let l = 0 , r = sClean.length
    for (let i =0; i < Math.floor(sClean.length/2); i++){
        if(sClean[l] != sClean[r-1]){
            return false
        }
        l++
        r--
        
    }
    return true
}   

let s = "A man, a plan, a canal: Panama"

console.log("answer to the first example is:",validPalindrome(s))