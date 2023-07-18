package sort

import "fmt"

func MergeSort(arr []int) {
	// Check Length is greater than 1
	output := make([]int, 0)
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := arr[:mid]
		right := arr[mid:]
		MergeSort(left)
		MergeSort(right)
		var i, j, k = 0, 0, 0
		for i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				output = append(output, left[i])
				i += 1
			} else {
				output = append(output, right[j])
				j += 1
			}
			k += 1
		}
		// Check if any element left in array
		for i < len(left) {
			output = append(output, left[i])
			i += 1
			k += 1
		}
		for j < len(right) {
			output = append(output, right[j])
			j += 1
			k += 1
		}
	}
	fmt.Println(output)
}

func CallMergeSort() {
	arr := []int{12, 11, 13, 5, 6, 7}
	MergeSort(arr)
}


func MergeSort1(arr []int)  {
	
}