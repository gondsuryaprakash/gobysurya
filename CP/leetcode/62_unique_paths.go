package leetcode

func UniquePaths(m int, n int) int {

	dp := [][]int{}
	for i := 0; i < m; i++ {

		arr := []int{}
		for j := 0; j < n; j++ {
			arr = append(arr, -1)
		}
		dp = append(dp, arr)

	}

	return CountPaths(0, 0, m, n, dp)

}

func CountPaths(i, j, n, m int, dp [][]int) int {
	if (i == n-1) && (j == m-1) {
		return 1
	}
	if i >= n || j >= m {
		return 0
	} else {
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		dp[i][j] = CountPaths(i+1, j, n, m, dp) + CountPaths(i, j+1, n, m, dp)
		return dp[i][j]
	}

}
