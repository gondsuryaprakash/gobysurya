package leetcode

import "sort"

func EraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	last := intervals[0][1]
	res := 1

	for i := 1; i < len(intervals); i++ {
		if last <= intervals[i][0] {
			last = intervals[i][1]
			res++
		}
	}
	return n - res
}
