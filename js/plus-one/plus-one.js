function plusOne(digits){
    let number = BigInt(digits.join(""))+BigInt(1) //BigInt to handle bigger arrays that will become bigger numbers
    let res = String(number).split("").map(e => Number(e))
    return res

}

let digits = [1,2,3]
console.log("answer to the first example is:",plusOne(digits))