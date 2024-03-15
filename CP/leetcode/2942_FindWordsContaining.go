package leetcode

import "strings"

func FindWordsContaining(words []string, x byte) []int {
	result := []int{}
	for i, v := range words {
		if strings.Contains(v, string(x)) {
			result = append(result, i)
		}
	}
	return result
}
