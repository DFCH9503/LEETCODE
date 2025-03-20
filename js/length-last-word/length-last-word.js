function lengthOfLastWord(s){
    let sSplit = s.split(" ").filter((word)=>word.length>=1)
    let idxLastWord = sSplit.length -1
    return sSplit[idxLastWord].length
}

let s = "Hello World"
console.log("answer to the first example is:",lengthOfLastWord(s))