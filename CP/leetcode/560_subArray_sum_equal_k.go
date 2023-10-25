package leetcode

import "sort"

func SubarraySum(nums []int, k int) int {

	count := 0
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		newSum := k - nums[i]
		if newSum == 0 {
			count++
			continue
		} else {
			j := i + 1
			jSum := 0
			for j < len(nums) {
				newSum = newSum - nums[j]
				jSum += nums[j]

				if newSum < 0 {
					j++
					break
				} else if newSum == 0 {

					count++
					newSum = newSum + jSum
					
				}
				j++

			}
		}

	}

	return count

}
