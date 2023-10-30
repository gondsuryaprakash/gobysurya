package leetcode

func MyPow(x float64, n int) float64 {
	var isNegative bool
	if n < 0 {
		isNegative = true
		n = -1 * n
	}

	var ans float64 = 1

	for n > 0 {
		if n%2 == 1 {
			ans = ans * x
			n = n - 1
		} else {
			x = x * x
			n = n / 2
		}
	}

	if isNegative {
		return 1 / ans
	}

	return ans
}
