package leetcode

import "strings"

func LengthOfLastWord(s string) int {

	s =strings.TrimSpace(s)
	sn:= strings.Split(s," ")
	return len(sn[len(sn)-1])

}
