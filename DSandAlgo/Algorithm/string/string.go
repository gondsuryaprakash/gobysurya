package string

// 12321   12321

func PalindromInt(n int) bool {
	rev := Reverse(n)

	return rev == n
}

func Reverse(n int) int {

	reverse := 0
	for n > 0 {
		rem := n % 10
		reverse = reverse*10 + rem
		n = n / 10
	}

	return reverse

}

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
