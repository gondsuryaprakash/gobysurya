package leetcode

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {

	candidates = []int{1, 1, 2, 5, 6, 7, 10}
	target = 8
	sort.Ints(candidates)
	// [10,1,2,7,6,1,5]
	// [1,1,2,5,6,7,10]

	start := 0
	lastIndex := len(candidates) - 1

	for start < lastIndex {

		nextStart := start + 1

		// Take sum of the element
		sum := start + nextStart + lastIndex
		if sum < target {
			for nextStart < lastIndex {
				nextStart += 1
				sum = start + nextStart + lastIndex
				if sum == target {
					
				}
				if sum < target {
					nextStart += 1
				}
			}

		} else if sum > target {
			lastIndex -= 1
		} else {

		}

	}

	return [][]int{}

}
