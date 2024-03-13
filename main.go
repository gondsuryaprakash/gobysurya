package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {
	//  data1 := [][]int{{1, 2}}
	// data := [][]int{{1, 2, 4}, {3, 3, 1}}

	// data1 := [][]int{{10}}
	// data := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	// data := [][]int{{1, 2}, {2, 3}}

	// data := [][]int{{1, 3}, {2, 3}, {3, 1}}

	// token := []int{2,1,3}
	// token := []int{5, 4, 3, 8}
	// token := []int{200, 100}
	// token := []int{91, 4, 75, 70, 66, 71, 91, 64, 37, 54}

	// score := 20
	// nums := []int{3, 4, 5, 2}
	// apple := []int{2, 83, 62}

	// order := "cba"
	// s := "abcd"

	// order := "kqep"
	// s := "pekeq"

	// k := 3
	// data := []int{10, 3, 8, 9, 4}
	// nums1 := []int{1, 1, 2}
	// nums2 := []int{1, 2, 3}
	// nums1 = []
	// nums2 = [1,2,3], k = 2
	// k := 2
	// result := leetcode.KSmallestPairs(nums1, nums2, k)
	result := leetcode.PivotInteger(1)
	fmt.Println(result)
}
