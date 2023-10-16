package leetcode

import (
	"fmt"
	"math"
)

func Trap(height []int) int {
	if len(height) == 1 {
		return 0
	}

	left := []int{}
	right := []int{}
	maxLeft := height[0]
	maxRight := height[len(height)-1]

	left = append(left, maxLeft)
	right = append(right, maxRight)

	for i := 1; i < len(height); i++ {
		if height[i] > maxLeft {
			maxLeft = height[i]
		}
		left = append(left, maxLeft)
	}

	for j := len(height) - 2; j > -1; j-- {
		if height[j] > maxRight {
			maxRight = height[j]
		}
		right = append(right, maxRight)
	}

	start:=0; 
	last:= len(right)-1
	for start < last {
		right[start], right[last] = right[last], right[start]
		start+=1
		last-=1
	}
	fmt.Println("Left :", left)
	fmt.Println("Right :", right)

	var totalWater int

	for i := 0; i < len(height); i++ {

		waterCaptured := int(math.Min(float64(left[i]), float64(right[i]))) - height[i]
		totalWater += waterCaptured

	}

	return totalWater

}
