package leetcode

func Intersection(nums1 []int, nums2 []int) []int {

	rc := make(map[int]struct{})

	for _, v := range nums1 {
		if _, ok := rc[v]; !ok {
			rc[v] = struct{}{}
		}
	}

	result := []int{}

	rcAdded := make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := rc[v]; ok {
			if _, ok2 := rcAdded[v]; !ok2 {
				rcAdded[v] = struct{}{}
				result = append(result, v)
			}
		}
	}
	return result
}
