package sort

func Swap(arr []int, x, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}

func InsertionSort(arr []int) []int {

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
