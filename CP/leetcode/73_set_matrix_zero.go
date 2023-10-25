package leetcode

import "fmt"

func SetZeroes(matrix [][]int) {

	rowArray := []int{}
	colArray := []int{}
	for i := 0; i < len(matrix); i++ {
		rowArray = append(rowArray, 0)
	}

	for j := 0; j < len(matrix[0]); j++ {
		colArray = append(colArray, 0)
	}

	for row := 0; row < len(matrix); row++ {

		for col := 0; col < len(matrix[0]); col++ {

			if matrix[row][col] == 0 {
				rowArray[row] = 1
				colArray[col] = 1
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if rowArray[row] == 1 || colArray[col] == 1 {
				matrix[row][col] = 0
			}
		}
	}

	fmt.Println(matrix)
}
