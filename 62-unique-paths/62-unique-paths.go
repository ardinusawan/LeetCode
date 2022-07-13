func uniquePaths(m int, n int) int {
    // backtracking
    // we find a way from finish back to start (reverse)
    // mark target as 1, then move backward to left and up
    // in new position, do calculate right + down value
    
    matrix := make([][]int, m)
    for c := 0; c < m; c++{
        matrix[c] = make([]int, n)
    }
    
    for row := m - 1; row >= 0; row-- {
        for col := n - 1; col >= 0; col-- {
            
            // mark target as 1
            if row == m - 1 && col == n - 1 {
                matrix[row][col] = 1
            } else {
                // use modules as trick if out of boundary value will become 0
                // initial value
                right := matrix[row][(col + 1) % n]
                down := matrix[(row + 1) % m][col]
                 
                // from current position, posibility of reaching target is 
                // by goes down or right. So every move we have sub-problem.
                // if we only have 2x2 box, the value will be like this
                // [2][1]
                // [1][1]
                matrix[row][col] = right + down    
            }
        }
    }
    return matrix[0][0]
}