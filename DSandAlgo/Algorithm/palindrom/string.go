package palindrom

func CheckPalindrom(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}
