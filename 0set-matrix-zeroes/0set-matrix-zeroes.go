func setZeroes(matrix [][]int)  {
    // flag = 0; means col/row need to be zeroed
    // we can use existing matrix to store flag, except [0][0] because it overlap
    // existing matrix mean the first row/col
    // rowZero = [0][0]
    
    ROWS, COLS := len(matrix), len(matrix[0])
    rowZero := false
    
    // iterate matrix to find 0
    // when found, mark parent row/col as 0
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if (matrix[r][c]) == 0 {
                matrix[0][c] = 0
                if (r > 0) {
                    matrix[r][0] = 0
                } else {
                    rowZero = true
                }
            }
        }
    }
    
    // make matrix 0 as per flag
    // from index 1, so we not remove flag value
    for r := 1; r < ROWS; r++ {
        for c := 1; c < COLS; c++ {
            if (matrix[0][c] == 0) || (matrix[r][0] == 0) {
                matrix[r][c] = 0
            }
        }
    }
    
    // make row and col in index 0 to zero
    if (matrix[0][0] == 0) {
        for i:=0; i<ROWS; i++ {
            matrix[i][0] = 0;
        }
    }
    
    if (rowZero) {
        for i:=0; i<COLS;i++ {
            matrix[0][i] = 0;
        }
    }
}