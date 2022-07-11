import (
    "container/list"
)

var (
    q *list.List
    visit [][]bool
)


func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    
    // rows, cols := len(grid), len(grid[0])
    islands := 0
    
    visit = make([][]bool, len(grid))
    for r := range grid {
        visit[r] = make([]bool, len(grid[0]))
    }
    
    q = list.New()
    
    for r := range grid {
        for c := range grid[0] {
            if string(grid[r][c]) == "1" && visit[r][c] == false {
                bfs(r, c, grid)
                islands += 1
            }
        }
    }
    
    return islands
}

func bfs(r int, c int, grid [][]byte) {
    q = list.New()
    visit[r][c] = true
    q.PushBack([]int{r, c})
    for q.Len() > 0 {
        row, col := q.Front(), q.Front()
        q.Remove(row)
        directions := make([][]int, 4)
        for d := range directions {
            directions[d] = make([]int, 2)
        }
        directions[0], directions[1], directions[2], directions[3] = []int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}
        for _, v := range directions {
            dr, dc := v[0], v[1]
            r, c = row.Value.([]int)[0] + dr, col.Value.([]int)[1] + dc
            
            if r >= 0 && r < len(grid) &&
            c >= 0 && c < len(grid[0]) &&
            string(grid[r][c]) == "1" &&
            visit[r][c] == false {
                q.PushBack([]int{r, c})     
                visit[r][c] = true
            }
        }
    }
}