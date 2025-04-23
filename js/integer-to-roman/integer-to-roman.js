function intToRoman(num){
    let numbers = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
    let strings = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]
    let res = ""


    while (num > 0){
        for (let i = 0; i < numbers.length; i++){
            if (num >= numbers[i]){
                res = res + strings[i]
                num = num - numbers[i]
                break
            }
        }
    }
    return res
}   

let num = 3479

console.log("answer to the first example is:",intToRoman(num))