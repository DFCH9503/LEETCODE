function mySqrt(x){
    if (x < 2) return x
    let root = x / 2
    while (true) {
        const approximateRoot = (root + x / root) / 2
        if (Math.floor(approximateRoot) === Math.floor(root)) {
            return Math.floor(approximateRoot)
        }
        root = approximateRoot
    }
}
let x=4
console.log("answer to the first example is:",mySqrt(4))