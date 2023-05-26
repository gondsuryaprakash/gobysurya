package leetcode

import "fmt"

func Shuffle(nums []int, n int) []int {
	mid := len(nums) / 2

	arr1 := nums[:mid]
	arr2 := nums[mid:]

	fmt.Println(arr1)
	fmt.Println(arr2)
	output := make([]int, len(nums))
	for i := 0; i < len(arr1); i++ {
		output = append(output, arr1[i], arr2[i])
	}
	fmt.Println(output)
	return output
}
