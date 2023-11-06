package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {

	// nums := []int{1, 3, 2, 3, 1}
	// nums2 := []int{2, 4, 3, 5, 1}

	// nums3 := []int{1, 0, -1, 0, -2, 2}
	// nums4 := []int{2, 2, 2, 2}
	// nums5 := []int{2, 2, 2, 2, 2}

	// nums6 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	nums7 := []int{100, 4, 200, 1, 3, 2}
	res := leetcode.LongestConsecutive(nums7)

	fmt.Println("Res: ", res)

}
