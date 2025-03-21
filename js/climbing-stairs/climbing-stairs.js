function climbStairs(n){
    let secondLast = 0
    let last = 1
    let temp = 0
    for (let i=1 ; i<=n ; i++){
        temp = secondLast + last
        secondLast = last
        last = temp
    }
    return last
}

let n =4
console.log("answer to the first example is:",climbStairs(n))
