function generateParenthesis(n){

    const res = []

    function dfs(openP, closeP, s) {
        if (openP === closeP && openP + closeP === n * 2) {
            res.push(s)
            return
        }

        if (openP < n) {
            dfs(openP + 1, closeP, s + "(")
        }

        if (closeP < openP) {
            dfs(openP, closeP + 1, s + ")")
        }
    }

    dfs(0, 0, "")

    return res   
}

let n = 3 
console.log(`With ${n} pairs of parentheses all the well formed couples are:`, generateParenthesis(n))