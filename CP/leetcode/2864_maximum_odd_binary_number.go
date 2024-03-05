package leetcode

import "strings"

func MaximumOddBinaryNumber(s string) string {
	numberOfOnes := strings.Count(s, "1")
	if numberOfOnes == 0 {
		return s
	}
	numberOfZeros := strings.Count(s, "0")
	if numberOfZeros == 0 {
		return s
	}

	result := ""

	for i := 0; i < numberOfOnes-1; i++ {
		result += "1"
	}
	for i := 0; i < numberOfZeros; i++ {
		result += "0"
	}

	result += "1"

	return result
}
