package leetcode

func MinIncrementForUnique(nums []int) int {

	minNumber := 0
	mp := make(map[int]int)

	for _, v := range nums {
		if v, ok := mp[v]; !ok {
			mp[v] = 0
		} else {
			mp[v] += 1
		}
	}
	for k, v := range mp {
		originalK := k
		if v > 0 {
			for {
				k += 1
				minNumber += 1
				if _, ok := mp[k]; !ok {
					mp[originalK] -= 1
					break
				}
			}

		}
	}
	return minNumber

}
