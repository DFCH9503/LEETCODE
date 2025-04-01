function longestPalindrome(s){
    let start = 0
    let end = 0
    if (s == null || s.length < 1) {
        return ""
    }
    for (let i = 0; i < s.length; i++){
        let len1 = expandAroundCenter(s, i, i)
        let len2 = expandAroundCenter(s, i, i+1)
        let maxLen = Math.max(len1, len2)
        if (maxLen > end-start){
            start = i - (maxLen-1) / 2
            end = i + maxLen/2
        }
    }
    let res=""
    s.length % 2 != 0 ? res = s.slice(start, end+1) : res = s.slice(start+1, end+1)
    return res 
}

function expandAroundCenter(s, left, right){
    while (left >= 0 && right < s.length && s.charAt(left) == s.charAt(right)){
        left--
        right++
    }
    return right-left-1
}

let s = "babad"
console.log("answer to the first example is:",longestPalindrome(s))