package leetcode

func MissingNumber(nums []int) int {
	result := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		result = result ^ nums[i] ^ i
	}
	return result ^ n
}
