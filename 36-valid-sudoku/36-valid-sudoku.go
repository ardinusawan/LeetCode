// https://leetcode.com/problems/valid-sudoku/
// remember: Only the filled cells need to be validated
// check row: create rowMap[byte]bool. iterate board[row][col], check if value exist in map. if no, fill map
// check row: create rowCol[byte]bool. same like check row but do in reverse
// create miniSudoku like this: map[1,1][5,3,7]
// map[1,1] = key = [row/3, col/3]. same like before, iterate over value, if not exist fill map

func isValidSudoku(board [][]byte) bool {
    miniSudoku := make(map[[2]int][]byte) // Map to store values in each 3x3 mini-grid

    // Iterate over each row in the board
    for row, _ := range board {
        rowMap := make(map[byte]bool) // Map to check for duplicate values in a row
        colMap := make(map[byte]bool) // Map to check for duplicate values in a column

        // Iterate over each column in the current row
        for col, _ := range board[row] {
            // check row: Create a map to track values in the current row
            // If the value already exists in the rowMap, there is a duplicate
            if rowMap[board[row][col]] {
                return false
            }

            // check column: Create a map to track values in the current column
            // If the value already exists in the colMap, there is a duplicate
            if colMap[board[col][row]] {
                return false
            }

            // check 3x3 (miniSudoku): Iterate over the values in each 3x3 mini-grid
            // If the value already exists in the miniSudoku map for the corresponding mini-grid, there is a duplicate
            for _, v := range miniSudoku[[2]int{row / 3, col / 3}] {
                if v == board[row][col] {
                    return false
                }
            }
            
            // Fill the maps: Add the current value to the respective rowMap, colMap, and miniSudoku map
            // Only fill the maps if the current cell contains a digit (ASCII code between 49 and 57)
            if (int(board[row][col]) >= 49 && int(board[row][col]) <= 57) {
                rowMap[board[row][col]] = true
                miniSudoku[[2]int{row / 3, col / 3}] = append(miniSudoku[[2]int{row / 3, col / 3}], board[row][col])
            }
            if (int(board[col][row]) >= 49 && int(board[col][row]) <= 57) {
                colMap[board[col][row]] = true
            }
        }
    }

    return true // The board is valid according to Sudoku rules
}
