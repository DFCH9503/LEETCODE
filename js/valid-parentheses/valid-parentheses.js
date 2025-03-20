function isValid(s){
    let stack = []
    let hash = {
        ")": "(",
        "]": "[",
        "}": "{",
    }
    for (i = 0 ; i < s.length ; i++){
        if (s[i] in hash){
            if(stack.length && stack[stack.length - 1] == hash[s[i]]){
                stack.pop()
            } else {
                return false
            } 
        } else {
            stack.push(s[i])
        }
    }
    return stack.length == 0

}

let s= "()"

console.log("answer to the first example is:",isValid(s))