package leetcode

func CountPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count += 1
			}
		}
	}
	return count
}

func isPrefixAndSuffix(str1, str2 string) bool {

	if len(str1) > len(str2) {
		return false
	}

	str2Prefix := str2[:len(str1)]
	str2Suffix := str2[len(str2)-len(str1):]

	if str2Prefix == str1 && str2Suffix == str1 {
		return true
	}

	return false
}
