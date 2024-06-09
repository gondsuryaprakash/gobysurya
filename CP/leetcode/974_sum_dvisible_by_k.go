package leetcode

func subarraysDivByK(nums []int, k int) int {
	cnt := make(map[int]int)
	prefixSum := 0
	cnt[0] = 1

	for _, num := range nums {
		prefixSum += num
		prefixSum = (prefixSum%k + k) % k
		cnt[prefixSum]++
	}

	res := 0
	for _, v := range cnt {
		res += v * (v - 1) / 2
	}

	return res
}
