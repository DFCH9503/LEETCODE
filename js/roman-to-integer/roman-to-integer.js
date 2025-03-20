const values = {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000
}

function romanToInt(s) {
    let sum = 0
    for (let i =0 ; i<s.length ; i++){
        if(i+1 < s.length && values[s[i]] < values[s[i+1]]){
            sum -= values[s[i]]
        }else{
            sum += values[s[i]]
        }
    }
    return sum
}


console.log("answer to the first example is:",romanToInt("MCMXCIV"))
