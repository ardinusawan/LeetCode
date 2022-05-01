func isValidSudoku(board [][]byte) bool {
    rows := make(map[int][]string, len(board))
    cols := make(map[int][]string, len(board[0]))
    squares := make(map[int]map[int][]string, len(board)/3) // to check every 3x3 box
    for i, _ := range board {
        squares[i] = make(map[int][]string, len(board[i])/3)
    }
    
    
    for r, _ := range board {
        for c, _ := range board[r] {
            v := string(board[r][c])
            
            if v == "." {
                continue
            }
            
            if contains(v, rows[r]) || // if value already in row
                contains(v, cols[c]) || // if value already in column
            contains(v, squares[r/3][c/3]) { // if value already in 3x3 box
                return false
            }
            
            rows[r] = append(rows[r], string(board[r][c]))
            cols[c] = append(cols[c], string(board[r][c]))
            squares[r/3][c/3] = append(squares[r/3][c/3], string(board[r][c]))    
            
        }
    }
    
    return true
}


func contains(s string, d []string) bool {
    for _, v := range d {
        if s == v {
            return true
        }
    }
    return false
}