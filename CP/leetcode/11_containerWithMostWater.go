package leetcode

import "math"


/*
Steps : 
1. Take two pointer left , right 
2. initialize maxArea with min value 
3. Run while loop til they are not met at one point 
2. Calculate the area between them
3. 


*/

func ContainerWithMostWater(arr []int) int {
	maxArea := math.MinInt
	left := 0
	right := len(arr) - 1
	for left < right {
		area := (right - left) * int(math.Min(float64(arr[left]), float64(arr[right])))
		maxArea = int(math.Max(float64(maxArea), float64(area)))
		if arr[left] < arr[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
