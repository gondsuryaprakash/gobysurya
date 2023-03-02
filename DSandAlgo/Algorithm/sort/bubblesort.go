package sort

func BubbleSort(arr []int) []int {
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
