package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	if target > matrix[len(matrix)-1][len(matrix[0])-1] && target < matrix[0][0] {
		return false
	}
	row := 0
	for row < len(matrix) {
		last := matrix[row][len(matrix[0])-1]
		if target == last {
			return true
		}
		if target < last {
			return binarySearch(matrix[row], target, 0, len(matrix[row])-1)
		}
		row++

	}
	return false
}

func binarySearch(arr []int, target, left, right int) bool {

	if right >= left {

		mid := left + (right-left)/2

		if arr[mid] == target {
			return true
		}

		if arr[mid] > target {
			return binarySearch(arr, target, left, mid-1)

		}
		return binarySearch(arr, target, mid+1, right)
	}

	return false
}
