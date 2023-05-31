package leetcode

import (
	"fmt"
	"sort"
)

func FourSum(arr []int, target int) [][]int {

	result := [][]int{}
	j := len(arr) - 1
	i := 0
	sort.Ints(arr)
	fmt.Println(arr)
	for i-1 < j-1 {
		k := i + 1
		l := j - 1

		for k < l {
			sum := arr[i] + arr[k] + arr[l] + arr[j]
			if sum == target {
				result = append(result, []int{arr[i], arr[k], arr[l], arr[j]})
				for arr[k] == arr[k+1] && k < l {
					k++
				}
			} else if sum > target {
				for arr[l] == arr[l-1] && l > k {
					l--
				}
			} else {
				for arr[k] == arr[k+1] && k < l {
					k++
				}
			}
		}

		i++
		j--
	}

	return result
}

func FourSum2(arr []int, target int) [][]int {
	result := [][]int{}
	n := len(arr)
	for i := 0; i < len(arr)-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < len(arr)-2; j++ {

			if j != i+1 && arr[j] == arr[j+1] {
				continue
			}

			k := j + 1
			l := n - 1

			for k < l {
				sum := arr[i] + arr[j] + arr[k] + arr[l]
				if sum == target {
					result = append(result, []int{arr[i], arr[j], arr[k], arr[l]})
					k++
					l--
					for k < l && arr[k] == arr[k+1] {
						k++
					}
					for k < l && arr[l] == arr[l-1] {
						l--
					}

				}
			}

		}
	}

	return result
}
