package leetcode

func VowelStrings(words []string, queries [][]int) []int {

	memeCached := make(map[int]bool)

	res := []int{}

	for _, query := range queries {
		queryCount := 0
		for i := query[0]; i <= query[1]; i++ {
			if isCache, ok := memeCached[i]; ok {
				if isCache {
					queryCount++
				}
			} else {
				if isEndIsStartAndEndWithVowel(words[i]) {
					queryCount++
					memeCached[i] = true
				} else {
					memeCached[i] = false
				}
			}

		}
		res = append(res, queryCount)
	}
	return res
}
func isVowel(str string) bool {
	return str == "a" || str == "e" || str == "i" || str == "o" || str == "u"

}
func VowelStrings2(words []string, queries [][]int) []int {

	m := len(words)
	count := 0
	prefixSum := make([]int, m)

	for i := 0; i < m; i++ {
		word := words[i]

		if isVowel(string(word[0])) && isVowel(string(word[len(word)-1])) {
			count++
		}
		prefixSum[i] = count
	}

	res := make([]int, len(queries))

	for i, query := range queries {

		l := query[0] - 1
		r := query[1]
		if l < 0 {
			res[i] = prefixSum[r]
		} else {
			res[i] = prefixSum[r] - prefixSum[l]
		}
	}
	return res
}

func isEndIsStartAndEndWithVowel(str string) bool {
	vowelMap := map[string]struct{}{
		"a": {},
		"e": {},
		"i": {},
		"o": {},
		"u": {},
	}
	if len(str) == 0 {
		if _, ok := vowelMap[str]; !ok {
			return false
		}
		return false
	}

	_, l1 := vowelMap[string(str[0])]
	_, l2 := vowelMap[string(str[len(str)-1])]

	if l1 && l2 {
		return true
	}
	return false
}
