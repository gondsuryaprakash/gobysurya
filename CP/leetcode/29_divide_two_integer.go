package leetcode

import "math"

func Divide(dividend int, divisor int) int {

	isNegative := false
	if dividend == 0 {
		return 0
	}
	if dividend < 0 && divisor < 0 {
		dividend = -(dividend)
		divisor = -(divisor)
	} else if dividend < 0 && divisor > 0 {
		dividend = -(dividend)
		isNegative = true
	} else if dividend > 0 && divisor < 0 {
		divisor = -(divisor)
		isNegative = true
	}

	if divisor == 1 && dividend > math.MaxInt32 {
		if isNegative {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	return 0

}
