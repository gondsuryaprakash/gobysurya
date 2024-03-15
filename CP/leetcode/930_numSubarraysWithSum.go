package leetcode

func NumSubarraysWithSum(nums []int, goal int) int {
	count := 0
	slidIngWindowMinLen := goal
	if goal == 0 {
		slidIngWindowMinLen = 1
	}
	for slidIngWindowMinLen < len(nums) {
		for i := 0; i < len(nums)-slidIngWindowMinLen+1; i++ {
			slidingSum := 0
			for j := i; j < i+slidIngWindowMinLen; j++ {
				slidingSum += nums[j]
			}
			if slidingSum == goal {
				count++
			}
		}
		slidIngWindowMinLen += 1
	}

	return count
}
