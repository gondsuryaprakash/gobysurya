package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func MyAtoi(s string) int {
	sArray := strings.Split(s, "")
	fmt.Println(sArray)
	mp := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

	num := ""
	var isNegative bool
	for i := 0; i < len(sArray); i++ {
		if sArray[i] == "-" {
			isNegative = true
		}
		if _, ok := mp[sArray[i]]; ok {

			num += sArray[i]
		}
	}
	if isNegative {
		num = "-" + num
	}
	fmt.Println(num)
	n, _ := strconv.Atoi(num)
	return n
}
