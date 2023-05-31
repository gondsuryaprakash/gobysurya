package leetcode

import "fmt"

func RemoveDuplicates(nums []int) int {
	k := 0
	arr := nums
	for i := 0; i < len(arr); i++ {
		j := i + 1
		for j < len(arr) {
			if arr[i] == arr[j] {
				k++
				nums = append(nums[:i], nums[i+1:]...)				
			} else {
				break
			}
			j++

		}
	}
	fmt.Println(nums)
	return k
}
