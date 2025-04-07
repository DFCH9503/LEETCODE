function myAtoi(s){
    s = s.trim()
    if (s.length == 0) return 0

    let sign = 1, i =0, res = 0
    if (s[i] == "-"){
        sign = -1
        i++
    }else if (s[i] == "+"){
        i++
    }

    while (i < s.length && s[i] >= "0" && s[i] <= "9"){
        res = res * 10 + (s[i]-"0")
        if(sign * res > 2**31-1) return 2**31-1
        if(sign * res < -(2**31)) return -(2**31)
        i++
    }
    return sign * res
}

let s = "4193 with words"
console.log("answer to the first example is:",myAtoi(s))