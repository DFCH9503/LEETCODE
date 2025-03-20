package main

import (
	"fmt"
	"strconv"
    "math" //comment for brute approach
)

func toStr(x int)string{
    xStr := strconv.Itoa(x)
    return xStr
}

//uncomment for brute approach 
// func reverse(str string) (result string) { 
//     for _, v := range str { 
//         result = string(v) + result 
//     } 
//     return
// } 


func isPalindrome(x int) bool {

    //brute approach
    // xStr := toStr(x)
	// xRev := reverse(xStr)
	// return xStr == xRev

    //using other approach, two pointer
    left , right := 0 , len(toStr(x))
    xString := toStr(x)
    xRune := []rune(xString)
    fmt.Println(xRune)
    for i := left; i< int(math.Floor(float64(right)/2)); i++ {
        charRight := xRune[right-(i+1)]
        charLeft := xRune[i]
        if charLeft == charRight{
            continue
        }else {
            return false
        }
    }
    return true
}


func main(){
    fmt.Println("answer to the first example is:",isPalindrome(121))
}


