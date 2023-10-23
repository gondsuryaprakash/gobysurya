package leetcode

func Rotate(matrix [][]int) {

	for col := 0; col < len(matrix[0]); col++ {
		for row := 0; row < len(matrix[0]); row++ {
			matrix[row][col] = 1
		}
	}

}
