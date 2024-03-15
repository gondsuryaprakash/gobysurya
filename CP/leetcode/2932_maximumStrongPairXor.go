package leetcode

import "math"

func MaximumStrongPairXor(nums []int) int {
	slideSize := 2
	maxResult := 0
	array := [][]int{}
	for i := 0; i < len(nums)-slideSize; i++ {
		pair := []int{}
		array = append(array, []int{i, i})
		for j := i; j < i+slideSize; j++ {
			pair = append(pair, nums[j])
		}
		
		if isStrongestPair(pair[0], pair[1]) {
			maxResult = int(math.Max(float64(maxResult), float64(pair[0]^pair[1])))

		}
	}
	return maxResult

}

func isStrongestPair(a, b int) bool {
	modDiff := a - b
	if modDiff < 0 {
		modDiff = -(modDiff)
	}
	var min int
	if a < b {
		min = a
	} else {
		min = b
	}
	return modDiff <= min
}
