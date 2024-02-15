package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {

	// data := []int{2, 2, 1, 1, 1, 2, 2}
	// data1 := []int{3, 1, -2, -5, 2, -4}
	// data2 := []int{1, 12, 1, 2, 5, 50, 3}
	data2 := []int{5, 5, 50}
	// data2 := []int{5, 5, 5}
	// data2 := []int{3, 2, 3}
	// data3 := []int{3}
	// grid := [][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}
	result := leetcode.LargestPerimeter(data2)
	fmt.Println(result)
}
