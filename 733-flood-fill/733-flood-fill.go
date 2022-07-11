var (
    img [][]int
    initColor int
    newColor int
    totalRow int
    totalCol int
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color {
        return image
    }
    
    img = image
    initColor = image[sr][sc]
    newColor = color
    
    totalRow = len(image)
    totalCol = len(image[0])
    paints(sr, sc)
    return img
}

// paints is dfs func to fill color 4-directionally
func paints(r, c int) {
    if img[r][c] == initColor {
        img[r][c] = newColor
        if r >= 1 {
            paints(r-1, c)
        }
        if r + 1 < totalRow {
            paints(r+1, c)
        }
        if c >= 1 {
            paints(r, c-1)
        }
        if c + 1 < totalCol {
            paints(r, c+1)
        }
    }
}