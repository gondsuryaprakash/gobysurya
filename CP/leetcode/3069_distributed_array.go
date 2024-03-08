package leetcode

func ResultArray(nums []int) []int {

	if len(nums) <= 2 {
		return nums
	}
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	for i := 2; i < len(nums); i++ {

		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	// merge the array
	arr1 = append(arr1, arr2...)
	return arr1
}
