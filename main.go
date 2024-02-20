package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {

	// data := []int{2, 2, 1, 1, 1, 2, 2}
	// data := []int{3, 1, -2, -5, 2, -4}
	// data := []int{1, 5, 1, 2, 3, 4, 10000}
	// data := []int{4, 2, 7, 6, 9, 14, 12}
	// data := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	// data := []int{5, 5, 5}
	// data := []int{3, 2, 3}
	// data3 := []int{3}
	// meeting := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	// meeting := [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}

	// meeting := [][]int{{0, 10}, {1, 2}, {12, 14}, {15, 16}}, 2
	// meeting := [][]int{{0, 10}, {1, 2}, {12, 20}, {15, 16}} //2

	meeting := [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}} //3
	// meeting := [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	// meeting := [][]int{{0,1}}

	result := leetcode.MostBooked(3, meeting)
	fmt.Println(result)
}
