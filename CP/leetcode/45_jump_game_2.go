package leetcode

import "math"

func Jump(num []int) int {
	if len(num) == 0 {
		return 0
	}
	res := 0
	r, l := 0, 0
	for r < len(num)-1 {
		farthest := 0

		for j := l; j < r+1; j++ {
			farthest = int(math.Max(float64(farthest), float64(j+num[j])))
		}
		l = r + 1
		r = farthest
		res += 1
	}

	return res
}
