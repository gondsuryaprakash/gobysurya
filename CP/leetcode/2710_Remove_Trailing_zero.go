package leetcode

func RemoveTrailingZero(str string) string {

	right := len(str) - 1

	for right >= 0 {
		if string(str[right]) == "0" {
			str = str[:right]
			right--
		} else {
			break
		}
	}
	return str
}
