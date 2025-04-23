package main

import(
	"fmt"
)

func intToRoman(num int) string{
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    strings := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""

	for num > 0{
		for idx, val := range numbers{
			if num >= val{
				res = res + strings[idx]
				num = num - val
				break
			}
		}
	}
	return res
}

func main(){
	num := 3479
	fmt.Printf("the number %d in roman is: %s \n",num ,intToRoman(num))
}