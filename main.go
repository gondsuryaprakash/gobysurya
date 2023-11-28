package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {

	// arr := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	// arr1 := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	arr2 := [][]int{{1, 2}, {1, 2}, {1, 2}}

	// arr2 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	outPut := leetcode.EraseOverlapIntervals(arr2)

	fmt.Println(outPut)

}
