function isValidSudoku(board){
    let rows = Array.from({ length: 9 }, () => new Set())
    let cols = Array.from({ length: 9 }, () => new Set())
    let boxes = Array.from({ length: 9 }, () => new Set())

    for (let r = 0; r < 9; r++) {
        for (let c = 0; c < 9; c++) {
            if (board[r][c] === '.') {
                continue
            }

            let value = board[r][c]
            let boxIndex = Math.floor(r / 3) * 3 + Math.floor(c / 3)

            if (rows[r].has(value) || cols[c].has(value) || boxes[boxIndex].has(value)) {
                return false
            }

            rows[r].add(value)
            cols[c].add(value)
            boxes[boxIndex].add(value)
        }
    }

    return true    
}

let board = [["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]


console.table(board)
console.log(`The sudoku board represented before is valid: `, isValidSudoku(board))