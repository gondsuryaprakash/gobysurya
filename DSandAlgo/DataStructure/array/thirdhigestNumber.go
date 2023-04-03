package array

import "math"

func FindThirdHighestNumber(arr []int) int {

	var first, second, third = math.Inf(-1), math.Inf(-1), math.Inf(-1)

	for i := 0; i < len(arr); i++ {

		if arr[i] > int(first) {
			third = second
			second = first
			first = float64(arr[i])
		} else if arr[i] > int(second) {
			third = second
			second = float64(arr[i])
		} else if arr[i] > int(third) {
			third = float64(arr[i])

		}
	}
	return int(third)
}
