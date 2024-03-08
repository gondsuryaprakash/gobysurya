package leetcode

func MaxFrequencyElements(nums []int) int {

	record := make(map[int]int)

	maxCount := 0
	for _, v := range nums {
		if _, ok := record[v]; !ok {
			record[v] = 1
			val, _ := record[v]
			if val > maxCount {
				maxCount = record[v]
			}
		} else {
			record[v] += 1
			val, _ := record[v]
			if val > maxCount {
				maxCount = record[v]
			}
		}
	}

	sum := 0
	for _, v := range record {
		if v == maxCount {
			sum += v
		}
	}

	return sum

}
