package main

import(
	"fmt"
)

func isValidSudoku(board [][]byte) bool{
	rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    boxes := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            val := board[r][c]
            if val == '.' {
                continue
            }

            if rows[r][val] {
                return false
            }
            rows[r][val] = true

            if cols[c][val] {
                return false
            }
            cols[c][val] = true

            boxIndex := (r / 3) * 3 + c / 3
            if boxes[boxIndex][val] {
                return false
            }
            boxes[boxIndex][val] = true
        }
    }

    return true
}

func main(){
	board := [][]string{{"5","3",".",".","7",".",".",".","."},
		{"6",".",".","1","9","5",".",".","."},
		{".","9","8",".",".",".",".","6","."},
		{"8",".",".",".","6",".",".",".","3"},
		{"4",".",".","8",".","3",".",".","1"},
		{"7",".",".",".","2",".",".",".","6"},
		{".","6",".",".",".",".","2","8","."},
		{".",".",".","4","1","9",".",".","5"},
		{".",".",".",".","8",".",".","7","9"}}

		boardByte := make([][]byte, len(board))

		for idx1, row := range board{
			boardByte[idx1] = make([]byte, 0, len(row))
			for _, col := range row{
				boardByte[idx1] = append(boardByte[idx1], []byte(col)...)
			}
		}

	fmt.Printf("The input board is a valid sudoku board: %v", isValidSudoku(boardByte))
}