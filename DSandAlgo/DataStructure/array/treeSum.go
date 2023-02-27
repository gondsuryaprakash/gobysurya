package array

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return all the (unique) triplets that sum up to 0.
// 2 pointers approach. O(N) time complexity.
func ThreeSum() {
	input := []int{-1, 0, 1, 2, -1, -4}

	outPut := threeSum(input)
	fmt.Println("outPut : ", outPut)

}

func threeSum(arr []int) [][]int {

	sort.Ints(arr)

	triplets := make([][]int, 0)

	for i := 0; i < len(arr)-2; i++ {

		if i > 0 && arr[i] == arr[i+1] {
			continue

		}

		j, k := i+1, len(arr)-1

		for j < k {
			sum := arr[i] + arr[j] + arr[k]

			if sum > 0 {
				k--
				continue
			}
			if sum < 0 {
				j++
				continue
			}
			triplets = append(triplets, []int{arr[i], arr[j], arr[k]})

			for j < k && arr[j] == arr[j+1] {
				j++
			}
			for j < k && arr[k] == arr[k-1] {
				k--
			}
			j, k = j+1, k-1
		}

	}
	return triplets
}
