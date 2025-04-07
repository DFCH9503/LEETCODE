package main

import (
	"fmt"
	"strings"
)

func strToByte(s string)[]byte{
	return []byte(s)
}

func myAtoi(s string)int{
	sSlice := strings.TrimSpace(s)
	if len(sSlice) == 0{
		return 0
	}

	sign := 1
	i := 0
	res := 0
	negative := "-"
	negativeByte := strToByte(negative)
	positive := "+"
	positiveByte := strToByte(positive)
	zero := "0"
	zeroByte := strToByte(zero)
	nine := "9"
	nineByte := strToByte(nine)

	if sSlice[i] == negativeByte[0]{
        sign = -1
        i++
    }else if sSlice[i] == positiveByte[0]{
        i++
    }


	for i < len(sSlice) && sSlice[i] >= zeroByte[0] && sSlice[i] <= nineByte[0]{
		res = res * 10 + int(sSlice[i]-zeroByte[0])
		if sign * res > (1<<31-1){ return 1<<31-1
		}
		if sign * res < -(1<<31){ 
			return -(1<<31)
		}
		i++
	}

	return res * sign
}

func main(){
	s := "+1" 

	fmt.Println("answer to the first example is:",myAtoi(s))

}