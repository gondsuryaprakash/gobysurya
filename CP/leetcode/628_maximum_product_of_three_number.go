package leetcode

import (
	"math"
	"sort"
)

func MaximumProduct(nums []int) int {

	sort.Ints(nums)

	//check all +ve and -ve

	// var isPositive, isNegative bool = true, true
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] < 0 {
	// 		isPositive = false
	// 		break
	// 	}
	// }
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] > 0 {
	// 		isNegative = false
	// 		break
	// 	}
	// }

	// if isNegative || isPositive {
	// 	return nums[len(nums)-2] * nums[len(nums)-3] * nums[len(nums)-1]
	// }

	// take first two
	firstTwoProduct := nums[0] * nums[1] * nums[len(nums)-1]
	lastTwo := nums[len(nums)-2] * nums[len(nums)-3] * nums[len(nums)-1]

	return int(math.Max(float64(firstTwoProduct), float64(lastTwo)))
}
