func rotate(matrix [][]int) {
	// have 4 pointer: left, right, top, bottom.
	// matrix[row][colomn]
	// if matrix size more than 2x2, need we need to solve it recursively by using left and right pointer. Increase left, decrease right so we will get diamond shape of matrix. Do it counterclockwise!
	// In doing counterclockwise, only need to store matrix[top][left] in temporary object
	// for every cycle ( 4 rotation ) left will increase, right will decrease, top increase, buttom decrease. Think like it diamond shape for second, third, ... etc iteration

	left, right := 0, len(matrix)-1 // colomn
	for left < right {
		top, buttom := left, right // row

		// iterate for every side of matrix
		for i := 0; i < (right - left); i++ {

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
