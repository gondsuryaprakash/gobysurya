package leetcode

import (
	"fmt"
	"sort"
)

func LargestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	fmt.Println(nums)
	lastIndex := len(nums) - 1
	for 1 < lastIndex {
		lastSum := sumUptoLast(lastIndex-1, nums)
		if lastSum > int64(nums[lastIndex]) {

			return lastSum + int64(nums[lastIndex])
		} else {
			lastIndex--
		}
	}
	return -1
}

func sumUptoLast(upto int, nums []int) int64 {
	sumUptoLast := 0
	for i := 0; i <= upto; i++ {
		sumUptoLast += nums[i]
	}
	return int64(sumUptoLast)
}
