package leetcode

func CountGoodSubstrings(s string) int {

	sum := 0
	start := 0
	end := 0
	if len(s) < 3 {
		return 0
	}

	charMap := make(map[byte]int)
	for end-start+1 <= 3 {
		charMap[s[end]] += 1
		end++
	}

	var isUniqeData = true
	for _, v := range charMap {
		if v > 1 {
			isUniqeData = false
			break
		}
	}
	if isUniqeData {
		sum += 1
	}

	for i := 3; i < len(s); i++ {
		charMap[s[i]]++
		charMap[s[i-3]]--

		if charMap[s[i-3]] == 0 {
			delete(charMap, s[i-3])
		}

		isUniqeData := true
		for _, v := range charMap {
			if v > 1 {
				isUniqeData = false
				break
			}
		}
		if isUniqeData {
			sum += 1
		}

	}
	return sum
}

func isUnique(s string) bool {
	return s[0] != s[1] && s[1] != s[2] && s[0] != s[2]
}
