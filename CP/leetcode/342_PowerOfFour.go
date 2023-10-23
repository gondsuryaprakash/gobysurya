package leetcode

func IsPowerOfFour(n int) bool {

	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n > 3 {
		return n%4 == 0 && IsPowerOfFour(n/4)
	}
	return false
}
