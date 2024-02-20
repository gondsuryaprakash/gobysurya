package recursion

func Reverse(l, r int, arr []int) {
	if r >= l {
		return
	}
	swap(arr[r], arr[l])
	Reverse(l+1, r-1, arr)
}

func swap(a, b int) {
	a, b = b, a
}

func CheckPalindrom(str string) bool {
	return recursiveCheck(0, len(str)-1, str)
}

func recursiveCheck(start, end int, str string) bool {
	if start >= len(str)/2 {
		return true
	}
	if str[start] != str[end] {
		return false
	}
	return recursiveCheck(start+1, end-1, str)

}
