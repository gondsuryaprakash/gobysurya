package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {
	//  data1 := [][]int{{1, 2}}
	//  data2 := [][]int{{1, 3}, {2, 3}}
	// data := [][]int{{1, 2}, {2, 3}}

	// data := [][]int{{1, 3}, {2, 3}, {3, 1}}

	// token := []int{2,1,3}
	// token := []int{5, 4, 3, 8}
	// token := []int{200, 100}
	// token := []int{91, 4, 75, 70, 66, 71, 91, 64, 37, 54}

	// score := 20
	nums := []int{6, 1, 3, 1, 1, 8, 9, 2}
	result := leetcode.IsPossibleToSplit(nums)
	fmt.Println(result)
}
