package main

import(
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := (dividend < 0) != (divisor < 0)

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	var start int
	switch {
	case dividend <= math.MaxUint8:
		start = 7
	case dividend <= math.MaxUint16:
		start = 15
	case dividend <= 16_777_215:
		start = 23
	default:
		start = 31
	}

	var quotient int
	for i := start; i >= 0 && dividend > 0; i-- {
		if divisor<<i <= dividend {
			dividend -= divisor << i
			quotient |= 1 << i
		}
	}

	if sign {
		return -quotient
	} else {
		return quotient
	}
}

func main(){
	dividend, divisor := 10, 3 
	
	fmt.Printf("%d divided by %d (truncated) is: %d", dividend, divisor, divide(dividend, divisor))
}