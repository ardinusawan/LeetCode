import (
    "fmt"
)

func rotate(matrix [][]int)  {
	// have 4 pointer: left, right, top, bottom
    // if matrix size more than 2x2, need we need to solve it recursively
    // for every cycle ( 4 rotation ) left will increase, right will decrease, top increase, buttom decrease
	
	left, right := 0, len(matrix)-1 // colomn
	for left < right {
		top, buttom := left, right // row

        // iterate for every side of matrix
		for i:=0;i<(right - left);i++ { 
			// replace backward
			// matrix[row][colomn]
            // remember: decrease/increase with i counterclockwise!
            // imagine diamond shape

            topLeft := matrix[top][left+i] // store first data temporary

			matrix[top][left+i] = matrix[buttom-i][left]

			matrix[buttom-i][left] = matrix[buttom][right-i]
            
			matrix[buttom][right-i] = matrix[top+i][right]
            
			matrix[top+i][right] = topLeft
		}
        
        // decrease box size
        left += 1
        right -= 1
	}
}
