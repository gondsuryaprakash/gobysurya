package leetcode

import (
	"math"
	"sort"
)

func BagOfTokensScore(tokens []int, power int) int {

	// Base Condition

	score := 0
	ans := 0
	sort.Ints(tokens)

	start := 0
	last := len(tokens) - 1

	for start <= last && (power >= tokens[start] || score > 0) {
		for start <= last && power >= tokens[start] {
			power -= tokens[start]
			start++
			score++
		}
		ans = int(math.Max(float64(ans), float64(score)))
		if start <= last && score > 0 {
			power += tokens[last]
			last--
			score--
		}

	}
	return ans

}
