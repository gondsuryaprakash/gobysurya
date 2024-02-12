package leetcode

func MajorityElement(nums []int) int {

	// using extra space

	record := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; !ok {
			record[nums[i]] = 1
		} else {
			record[nums[i]] += 1
		}
	}
	maxValue := 0
	max := 0
	for k, v := range record {
		if v > max {
			max = v
			maxValue = k
		}
	}
	if max > len(nums)/2 {
		return maxValue
	}
	return 0
}
