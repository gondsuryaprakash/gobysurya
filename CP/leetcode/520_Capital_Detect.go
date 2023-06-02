package leetcode

import (
	"strings"
)

func DetectCapitalUse(word string) bool {
	w := word
	if word == strings.ToUpper(w) || word == strings.ToLower(w) || word == strings.ToUpper(string(w[0]))+strings.ToLower(w[1:]) {
		return true
	}

	return false
}
