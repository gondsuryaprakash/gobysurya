package leetcode

import "fmt"

func ReversePairs(nums []int) int {
	// Common Approach //It will fail for the long array
	count := 0
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for j < len(nums) {
			if nums[i] > 2*nums[j] {
				fmt.Printf("nums[%v] = %v , nums[%v] = %v ", i, nums[i], j, 2*nums[j])
				count++
			}
			j++
		}
	}
	return count
}

func ReversePairsUsingMergeSort(nums []int) int {
	return 0
}

func mergeSort(arr []int) {
	// Check Length is greater than 1
	output := make([]int, 0)
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := arr[:mid]
		right := arr[mid:]
		mergeSort(left)
		mergeSort(right)
		var i, j, k = 0, 0, 0
		for i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				output = append(output, left[i])
				i += 1
			} else {
				output = append(output, right[j])
				j += 1
			}
			k += 1
		}
		// Check if any element left in array
		for i < len(left) {
			output = append(output, left[i])
			i += 1
			k += 1
		}
		for j < len(right) {
			output = append(output, right[j])
			j += 1
			k += 1
		}
	}
	fmt.Println(output)
}
