package string

func LongestSubString(s string) int {

	start := 0
	currentSubStringLength := 0
	longestSubStringLength := 0

	record := make(map[rune]int)

	for index, char := range s {

		lastIndex, ok := record[char]

		if !ok || lastIndex < index-currentSubStringLength {
			currentSubStringLength++
		} else {
			if currentSubStringLength > longestSubStringLength {
				longestSubStringLength = currentSubStringLength
			}
			start = index + 1
			currentSubStringLength = index - start + 1
		}
		record[char] = index
	}

	if currentSubStringLength > longestSubStringLength {
		longestSubStringLength = currentSubStringLength
	}
	return longestSubStringLength
}
