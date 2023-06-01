package leetcode

func MergeAlternately(word1 string, word2 string) string {
	res := ""
	var i int
	var minLen int
	if len(word1) > len(word2) {
		minLen = len(word2)
	} else {
		minLen = len(word1)
	}
	for ; i < minLen; i++ {
		res += string(word1[i]) + string(word2[i])
	}
	if len(word1) > len(word2) {
		res += word1[i:]
	} else {
		res += word2[i:]
	}

	return res
}
