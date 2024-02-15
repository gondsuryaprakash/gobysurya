package leetcode

import "fmt"

func firstPalindrome(words []string) string {
	for _, v := range words {
		if isPalindrom(v) {
			return v
		}
	}
	return ""
}

func isPalindrom(str string) bool {
	fmt.Println(str)
	start := 0
	end := len(str) - 1
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}

	return true
}
