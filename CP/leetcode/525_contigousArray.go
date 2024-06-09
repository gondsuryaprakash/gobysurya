package leetcode

func FindMaxLength(nums []int) int {

	answer := 0

	for k := 2; k < len(nums); k++ {
		rc := make(map[int]int)
		start := 0
		end := 0

		for end-start+1 <= k {
			rc[nums[end]]++
			end++
		}
		if rc[0] == rc[1] {
			answer = max(answer, rc[0]+rc[1])
		}

		for i := k; i < len(nums); i++ {
			rc[nums[i]]++
			rc[nums[i-k]]--
			if rc[nums[i]] == rc[nums[i-k]] {
				answer = max(answer, rc[nums[i]]+rc[nums[i-k]])
			}
			if rc[nums[i-k]] == 0 {
				delete(rc, nums[i-k])
			}
		}
	}

	return answer
}	

func max(a, b int) int {
	if a > b {
		return a
	}

	return b

}
