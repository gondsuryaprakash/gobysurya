package array

func TwoSum(arr []int, target int) []int {
	record := make(map[int]int)
	for i, j := range arr {
		complement := target - j
		if val, ok := record[complement]; ok {
			return []int{val, i}
		}
		record[j] = i
	}
	return []int{}

}
