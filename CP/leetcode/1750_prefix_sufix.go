package leetcode

/*
1.Pick a non-empty prefix from the string s where all the characters in the prefix are equal
2.Pick a non-empty suffix from the string s where all the characters in this suffix are equal.
3.The prefix and the suffix should not intersect at any index.
4.The characters from the prefix and suffix must be the same.
5.Delete both the prefix and the suffix.
*/
func MinimumLength(s string) int {

	if len(s) == 1 {
		return 1
	}

	start := 0
	last := len(s) - 1
	length := len(s)
	for start < last {
		prefixLength := 1
		suffixLength := 1
		i := start
		j := last
		prefix := s[start]
		suffix := s[last]
		if prefix != suffix {
			return length
		}
		for i < last {
			if s[i] == s[i+1] {
				if i+1 != last {
					prefixLength += 1
				}
				i++
			} else {
				break
			}
		}
		for j > i && s[j] == prefix {
			if s[j] == s[j-1] {
				if j-1 != i {
					suffixLength += 1
				}
				j--
			} else {
				break
			}
		}
		if prefix == suffix {
			length = length - prefixLength
			length = length - suffixLength
			// remove string
			start = start + prefixLength
			last = last - suffixLength
		}
	}
	return length
}
