package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {
	// heights := []int{3, 4, 2, 5, 2}
	// results := leetcode.HeightChecker(heights)
	// fmt.Println(results)

	/*
		[4,9,5], nums2 = [9,4,9,8,4]

	*/
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}

	// result := leetcode.Intersect(nums1, nums2)

	// fmt.Println(result)

	// nums := []int{2, 5, 3, 4, 1}
	nums3 := []int{1, 2, 3, 4}
	// nums2 := []int{2, 1, 3}
	count := leetcode.NumTeams(nums3)
	fmt.Println(count)
}
