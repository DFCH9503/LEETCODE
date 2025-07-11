function maxProfit(prices){
    let minPrice = prices[0], maxPrice = prices[0], res = 0

    for (price of prices){
        if (price < minPrice){
			minPrice = price
			maxPrice = price
		}else if (maxPrice < price){
			maxPrice = price
		}
		if ((maxPrice - minPrice) > res){
			res = maxPrice - minPrice
		}
	}
	return res
}


let prices = [7, 1, 5, 3, 6, 4]
console.log(`The max profit for ${prices} is ${maxProfit(prices)}`)