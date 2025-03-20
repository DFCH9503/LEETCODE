function longestCommonPrefix(strs){
    if (strs.length == 0) {
        return ""
    }
    strs.sort()
    let first = strs[0], last = strs[strs.length-1]
    let result = ""
    for (i = 0 ; i < first.length ; i++){
        if (i < last.length && first[i] == last[i]){
            result += first[i].toString()
        }else{
            break
        }
    }
    return result
}

let array = ["flow", "flower", "flight"]
console.log("answer to the first example is:",longestCommonPrefix(array))