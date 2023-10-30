package leetcode

import "sort"

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	right := 0

	for right < len(nums2) {
		nums1[m+right] = nums2[right]
		right++

	}
	sort.Ints(nums1)

}
