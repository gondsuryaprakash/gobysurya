package contest

import "strings"

func IsSubstringPresent_v1(s string) bool {

	mp := make(map[byte]int)

	start := 0
	end := 0
	for end-start+1 <= 2 {
		mp[s[end]]++
		end++
	}

	result := ""
	for k, v := range mp {
		for v > 0 {
			result += string(k)
			v--
		}

	}

	reversedString := reverse(s)
	if strings.Contains(reversedString, result) {
		return true
	}

	for i := end; i < len(s); i++ {

		mp[s[i]]++
		mp[s[i-2]]--

		if mp[s[i-2]] == 0 {
			delete(mp, s[i-2])
		}
		subString := ""
		for k, v := range mp {
			for v > 0 {
				subString += string(k)
				v--
			}
		}
		if strings.Contains(reversedString, subString) {
			return true
		}

	}

	return false

}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func IsSubstringPresent_v2(s string) bool {

	start := 0
	end := 1
	reversedString := reverse(s)

	for end < len(s) {

		a := string(s[start])
		b := string(s[end])
		c := a + b

		if strings.Contains(reversedString, c) {
			return true
		}
		start++
		end++
	}
	return false
}
