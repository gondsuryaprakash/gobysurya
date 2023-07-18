package leetcode

import "math"

/*
arr := [2,3,1,1,4]
*/
func CanJump(nums []int) bool {
	var i int
	for reach := 0; i < len(nums) && i <= reach; i++ {
		reach = int(math.Max(float64(reach), float64(i+nums[i])))
	}
	return i == len(nums)

}
