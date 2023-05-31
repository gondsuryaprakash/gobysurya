package main

import (
	"fmt"

	"gotutorial.com/CP/leetcode"
)

func main() {
	rect := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := leetcode.ContainerWithMostWater(rect)

	fmt.Println(area)
}
