package leetcode

import "math"

func FindMaxConsecutiveOnes(nums []int) int {
	count := 0
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			c++
		} else {
			c = 0
		}
		count = int(math.Max(float64(c), float64(count)))
	}
	return count
}
