package leetcode

func IsPowerOfThree(n int) bool {

	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	if n > 1 {
		return n%3 == 0 && IsPowerOfThree(n/3)
	}

	return false
}
