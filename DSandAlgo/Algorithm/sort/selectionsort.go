package sort

func Swap(arr []int, x, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}

func SelectionSort(arr []int) []int {
	arrayLenght := len(arr)
	for i := 0; i < arrayLenght-1; i++ {
		min_Index := i
		for j := i + 1; j < arrayLenght; j++ {

			if arr[j] < arr[min_Index] {
				min_Index = j
			}
		}
		Swap(arr, i, min_Index)
	}
	return arr
}

func seletionSort(arr []int) []int {
	arrLength := len(arr)
	for i := 0; i < len(arr); i++ {
		min_index := i
		for j := i + 1; j < arrLength; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		Swap(arr, i, min_index)
	}
	return arr
}
