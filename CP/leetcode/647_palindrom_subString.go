package leetcode

func CountSubstrings(s string) int {

	count := 0

	// create possible subString
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			if checkPalindrom(s[i:j]) {
				count++
			}
		}
	}

	return count
}

func checkPalindrom(s string) bool {

	start := 0
	end := len(s) - 1
	for start < end {

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
