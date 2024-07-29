package leetcode

import "sort"

func HeightChecker(heights []int) int {

	v1 := make([]int, len(heights))
	copy(v1, heights)
	sort.Ints(v1)
	count := 0
	for i, v := range heights {
		if v != v1[i] {
			count += 1
		}
	}
	return count
}
