package leetcode

import (
	"math"
	"sort"
)

func FindMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 0

	last := -math.MaxInt

	for i := 0; i < len(points); i++ {
		if last < points[i][0] {
			last = points[i][1]
			res++
		}
	}

	return res
}
