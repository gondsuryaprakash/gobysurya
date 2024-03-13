package leetcode

import "math"

func PivotInteger(n int) int {
	prefixArr := []int{1}
	for i := 1; i < n; i++ {
		prefix := prefixArr[i-1] + i + 1
		prefixArr = append(prefixArr, prefix)
	}
	for i := 1; i <= n; i++ {
		sumUpToi := prefixArr[i-1]
		for j := i; j <= n; j++ {
			sumUpToi = sumUpToi - j
			if sumUpToi < 0 {
				break
			}
		}
		if sumUpToi == 0 {
			return i
		}
	}

	return -1
}

func PivoteIntegerV2(n int) int {
	sqrt := math.Sqrt(float64(n * (n + 1) / 2))

	if int(sqrt)%1 != 0 {
		return -1
	}
	return int(sqrt)
}
