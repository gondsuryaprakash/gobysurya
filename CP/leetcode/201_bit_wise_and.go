package leetcode

import "fmt"

func RangeBitwiseAnd(left int, right int) int {

	// bitwise and means that

	result := 0
	for i := left; i <= right; i++ {
		shift := i << 1
		fmt.Println(shift)
	}

	fmt.Println("result", result)

	return result
}
