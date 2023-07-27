package leetcode

import "strconv"

func IsValidSudoku(board [][]byte) bool {

	// check for the row
	record := make(map[string]bool)
	for rows := 0; rows < 9; rows++ {
		for cols := 0; cols < 9; cols++ {
			if string(board[rows][cols]) != "." {
				grid := (rows/3)*3 + cols/3
				rowKeys := "row" + string(board[rows][cols]) + strconv.Itoa(rows)
				colKeys := "cols" + string(board[rows][cols]) + strconv.Itoa(cols)
				gridKyes := "grid" + string(board[rows][cols]) + strconv.Itoa(grid)

				_, rowOk := record[rowKeys]
				_, colsOk := record[colKeys]
				_, gridOk := record[gridKyes]
				if rowOk || colsOk || gridOk {
					return false
				} else {
					record[rowKeys] = true
					record[colKeys] = true
					record[gridKyes] = true
				}
			}

		}
	}

	return true
}
