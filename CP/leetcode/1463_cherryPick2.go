package leetcode

// func CherryPickup(grid [][]int) int {
// 	rows := len(grid)
// 	cols := len(grid[0])

// 	dp := make([][][]int, rows)
// 	for i := range dp {
// 		dp[i] = make([][]int, cols)
// 		for j := range dp[i] {
// 			dp[i][j] = make([]int, cols)
// 			for k := range dp[i][j] {
// 				dp[i][j][k] = -1
// 			}
// 		}
// 	}

// 	return dfs(0, 0, cols-1, grid, dp)

// }

// func dfs(rows, cols1, cols2 int, grid [][]int, dp [][][]int) int {
// 	if rows >= len(grid) || cols1 < 0 || cols1 >= len(grid[0]) || cols2 < 0 || cols2 >= len(grid[0]) {
// 		return 0
// 	}

// 	if dp[rows][cols1][cols2] != -1 {
// 		return dp[rows][cols1][cols2]
// 	}

// 	sum := 0
// 	if rows < len(grid) {
// 		sum += grid[rows][cols1]

// 		if cols1 != cols2 {
// 			sum += grid[rows][cols2]
// 		}

// 		max := 0
// 		for i := cols1 - 1; i < cols1+1; i++ {
// 			for j := cols2 - 1; j < cols2+1; j++ {
// 				max = Max(max, dfs(rows+1, i, j, grid, dp))
// 			}
// 		}

// 		sum += max
// 		dp[rows][cols1][cols2] = sum
// 	}
// 	return sum
// }

func CherryPickup(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	// Initialize the dp array
	dp := make([][][]int, row)
	for i := range dp {
		dp[i] = make([][]int, col)
		for j := range dp[i] {
			dp[i][j] = make([]int, col)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	return dfs(0, 0, col-1, dp, grid)
}

func dfs(r, c1, c2 int, dp [][][]int, grid [][]int) int {
	if r >= len(grid) || c1 < 0 || c1 >= len(grid[0]) || c2 < 0 || c2 >= len(grid[0]) {
		return 0
	}

	if dp[r][c1][c2] != -1 {
		return dp[r][c1][c2]
	}

	result := 0
	if r < len(grid) {
		result += grid[r][c1]
		if c1 != c2 {
			result += grid[r][c2]
		}

		max := 0
		for i := c1 - 1; i <= c1+1; i++ {
			for j := c2 - 1; j <= c2+1; j++ {
				max = Max(max, dfs(r+1, i, j, dp, grid))
			}
		}
		result += max
		dp[r][c1][c2] = result
	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
