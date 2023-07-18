package sort

import "fmt"

func bubbleSort(arr []int) []int {
	arrayLength := len(arr)
	for i := 0; i < arrayLength-1; i++ {
		for j := 0; j < arrayLength-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
			}
		}
	}

	return arr

}

func bubbleSortImp2(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}

func CallBubbleSort() {
	arr := []int{12, 11, 13, 5, 6, 7}

	sortedArr := bubbleSortImp2(arr)
	sirtedArr2 := bubbleSort(arr)
	fmt.Println("SortedArr : ", sortedArr)
	fmt.Println("sirtedArr2 : ", sirtedArr2)
}
