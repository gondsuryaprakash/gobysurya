package leetcode

import "github.com/spf13/cast"

func CountAndSay(n int) string {

	// Base condition
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	result := "11"
	for i := 3; i <= n; i++ {
		c := 1
		result += "&"
		t := ""
		for j := 1; j < len(result); j++ {
			if result[j] != result[j-1] {
				t = t + cast.ToString(c)
				t = t + string(result[j-1])
				c = 1
			} else {
				c++
			}
		}
		result = t
	}
	return result

}
