package leetcode

func LargestDivisibleSubset(nums []int) []int {

	// BruteForce Approach
	result := []int{}
	for i := 0; i < len(nums); i++ {
		tempResult := []int{}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]%nums[j] == 0 || nums[j]%nums[i] == 0 {
				tempResult = append(tempResult, nums[i], nums[j])
			}
		}
		if len(tempResult) > len(result) {
			result = tempResult
		}
	}
	return result
}
