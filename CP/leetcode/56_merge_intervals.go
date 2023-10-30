package leetcode

import (
	"math"
	"sort"
)

func MergeIntervals(intervals [][]int) [][]int {

	ans := [][]int{}

	// Sort the array
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])

	for i := 1; i < len(intervals); i++ {

		current := intervals[i]
		last := ans[len(ans)-1]

		if current[0] <= last[1] {
			last[1] = int(math.Max(float64(last[1]), float64(current[1])))
		} else {
			ans = append(ans, current)
		}

	}

	return ans

}
