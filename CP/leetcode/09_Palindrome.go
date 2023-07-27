package leetcode

func IsPalindrome(x int) bool {
	var reverse int
	xCopy := x
	for xCopy > 0 {
		reverse = reverse*10 + xCopy%10
		xCopy = xCopy / 10
	}

	return reverse == x
}
