package leetcode

func Intersect(nums1 []int, nums2 []int) []int {

	result := []int{}
	mp := make(map[int]int)

	for _, v := range nums1 {
		mp[v] += 1
	}

	for _, v := range nums2 {
		if val, ok := mp[v]; ok && val > 0 {
			result = append(result, v)
			mp[v] -= 1
		}
	}

	return result
}
