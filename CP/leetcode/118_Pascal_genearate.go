package leetcode

import "fmt"

func Generate(numRows int) [][]int {

	numbers := [][]int{}

	for i := 0; i < numRows; i++ {
		newArray := []int{}
		for j := 0; j <= i; j++ {
			newArray = append(newArray, 0)
		}
		numbers = append(numbers, newArray)
	}

	for line := 0; line < numRows; line++ {
		for i := 0; i <= line; i++ {
			if line == i || i == 0 {
				numbers[line][i] = 1
			} else {	
				numbers[line][i] = numbers[line-1][i-1] + numbers[line-1][i]
			}
		}
	}

	fmt.Println(numbers)
	return numbers

}
