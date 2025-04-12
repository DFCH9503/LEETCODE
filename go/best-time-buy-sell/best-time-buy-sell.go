package main

import (
	"fmt"
)

func maxProfit(prices []int) int{
	minPrice := prices[0]
	maxPrice := prices[0]
	res := 0

	for _, price := range prices{
		if price < minPrice{
			minPrice = price
			maxPrice = price
		}else if maxPrice < price{
			maxPrice = price
		}
		if (maxPrice - minPrice) > res{
			res = maxPrice - minPrice
		}
	}
	return res
}

func main(){
	prices := []int{7,1,5,3,6,4}
	fmt.Printf("The max profit for %v is %d", prices, maxProfit(prices))
}