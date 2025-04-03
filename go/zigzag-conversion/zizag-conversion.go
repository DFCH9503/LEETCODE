package main

import(
	"fmt"
	"strings"
)

func createRows(numRows int) [][]string{
	rows := make([][]string, numRows)
	for i := range rows {
		rows[i] = []string{}
	}
	return rows
}


func zigzagConversion(s string, numRows int)string{
	if numRows == 1 || numRows >= len(s){
        return s
    }

	idx := 0
	d := 1
	rows := createRows(numRows)

	for _,char := range s{
		rows[idx] = append(rows[idx], string(char))
		if idx == 0 {
            d = 1
        } else if idx == numRows - 1 {
            d = -1
        }
        idx += d
	}
	res := ""
	for _, val := range rows{
		res += strings.Join(val,"")
	}

	return res
}	

func main(){
	s := "PAYPALISHIRING"
	numRows := 5

	fmt.Println("answer to the first example is:",zigzagConversion(s, numRows))
}