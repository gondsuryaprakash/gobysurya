package leetcode

import (
	"math"
)

func LongestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	hashMap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; !ok {
			hashMap[nums[i]] = struct{}{}
		}
	}
	maxSequence := -math.MaxInt32
	for val, _ := range hashMap {
		if _, ok := hashMap[val-1]; !ok {
			count := 1
			k := val
			for found(k, hashMap) {
				count += 1
				k += 1
			}
			maxSequence = int(math.Max(float64(count), float64(maxSequence)))
		}
	}
	return maxSequence - 1

}

func found(x int, hasmMap map[int]struct{}) bool {

	if _, ok := hasmMap[x]; ok {

		return ok

	}
	return false
}
