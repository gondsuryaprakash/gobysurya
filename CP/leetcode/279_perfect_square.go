package leetcode

import "math"

func NumSquares(n int) int {

	dp := []int{}
	for i := 0; i < n+1; i++ {
		dp = append(dp, 0)
	}

	for i := 0; i < n+1; i++ {

		dp[i] = i
		for j := 1; j < int(math.Sqrt(float64(i))+1); j++ {

			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-j*j])))

		}
	}

	return dp[n]
}
