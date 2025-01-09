package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {
	// words := []string{"aba", "bcb", "ece", "aa", "e"}
	// queries := [][]int{{0, 2}, {1, 4}, {1, 1}}
	words := []string{"a", "aba", "ababa", "aa"}

	ans := leetcode.CountPrefixSuffixPairs(words)
	fmt.Println(ans)
}
