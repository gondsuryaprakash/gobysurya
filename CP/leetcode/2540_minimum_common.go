package leetcode

import (
	"fmt"
	"sort"
)

func GetCommon(nums1 []int, nums2 []int) int {

	// Using the map
	rc := make(map[int]struct{})

	for _, v := range nums2 {
		rc[v] = struct{}{}
	}

	for i := 0; i < len(nums1); i++ {
		if _, ok := rc[nums1[i]]; ok {
			return nums1[i]
		}
	}
	return -1
}

// Using two Pointers
func GetCommon_v2(nums1 []int, nums2 []int) int {

	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}

		if nums1[i] < nums2[j] {
			i++
		}
		if nums1[i] > nums2[j] {
			j++
		}

	}
	return -1
}

func MinimumBoxes(apple []int, capacity []int) int {
	totalApple := 0

	for _, v := range apple {
		totalApple += v
	}

	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})

	fmt.Println(capacity)

	// return minimum number of box required

	capacitySum := 0

	for i := 0; i < len(capacity); i++ {

		capacitySum += capacity[i]

		if capacitySum >= totalApple {
			return i + 1
		}
	}
	return 0
}

func MaximumHappinessSum(happiness []int, k int) int64 {

	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	happinessSum := 0
	start := 0

	startDecrese := 0
	for start < k {

		if happiness[start]-startDecrese > 0 {
			happinessSum += happiness[start] - startDecrese
		}
		startDecrese += 1

		// for i := start + 1; i < len(happiness); i++ {
		// 	// decrement by 1
		// 	if happiness[i] > 0 {
		// 		happiness[i] = happiness[i] - 1
		// 	}
		// }
		start++
	}
	return int64(happinessSum)
}
