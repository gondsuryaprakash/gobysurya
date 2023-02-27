package array

import "fmt"

func MoveAllZeroAtEnd() {

	arr := []int{-1, 2, -3, 4, 5, 6, -7, 8, 9}

	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			if i != j {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
			j++
		}
	}

	fmt.Println(arr)
}
