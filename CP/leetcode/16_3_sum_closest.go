package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {

	// var closetSum int
	sort.Ints(nums)
	var start int
	var closestSum int
	end := len(nums)
	for start < end {
		sum := nums[start] + nums[start+1] + nums[end]
		if sum == target {
			return sum
		} else if sum > target {

		}

	}
	return closestSum

}
