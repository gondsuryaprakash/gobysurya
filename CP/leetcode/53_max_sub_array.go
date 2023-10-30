package leetcode

import "math"

func MaxSubArray(nums []int) int {
	maxSum := -math.MaxInt - 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum

}
