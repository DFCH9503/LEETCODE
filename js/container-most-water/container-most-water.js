function getMin(a, b){
    if (a < b) {
        return a
    }else{
        return b
    }
}

function getMax(a, b){
    if (a < b) {
        return b
    }else{
        return a
    }
}

function maxArea(height){
    let l = 0
	let r = height.length -1
	let res = 0

	while (l < r) {
        let currWidth = r-l
		let currHeight = getMin(height[l], height[r])
		res = getMax(res, currWidth*currHeight)
		if (height[l] <= height[r]) {
            l++
        } else {
            r--
        }
	} 
	return res
}   

let height = [1,8,6,2,5,4,8,3,7]

console.log("answer to the first example is:",maxArea(height))