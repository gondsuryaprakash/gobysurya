package leetcode

func Rotate(matrix [][]int) {

	transPoseArray := [][]int{}
	for i := 0; i < len(matrix); i++ {
		arr := []int{}
		for j := 0; j < len(matrix[0]); j++ {
			arr = append(arr, 0)
		}
		transPoseArray = append(transPoseArray, arr)
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			transPoseArray[row][col] = matrix[col][row]
		}
	}

	for j := 0; j < len(matrix); j++ {
		start := 0
		lastColum := len(matrix[0]) - 1
		for start < lastColum {
			transPoseArray[j][start], transPoseArray[j][lastColum] = transPoseArray[j][lastColum], transPoseArray[j][start]
			start++
			lastColum--
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = transPoseArray[i][j]
		}
	}

}
