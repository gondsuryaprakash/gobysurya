package leetcode

func IsPossibleToSplit(nums []int) bool {
	record := make(map[int]int)

	for _, v := range nums {
		if _, ok := record[v]; !ok {
			record[v] = 1
		} else {
			if record[v] >= 2 {
				return false
			}
			record[v] += 1
		}
	}

	return true
}
