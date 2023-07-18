package leetcode

func SearchRange(nums []int, target int) []int {
	return []int{}
}

func BinarySearch(arr []int, start, last, target int) (result []int) {

	for start <= last {

		mid := start + (last-start)/2

		if arr[mid] == target {
			result = append(result, mid)
			if last-mid >=1 {
				if arr[mid+1] == target {
					result = append(result, mid+1)
				}
			}
		}

		if arr[mid] < target {
			start = mid + 1
		} else {
			last = mid - 1
		}

	}
	if len(result) == 0 {
		result = []int{-1, -1}
	}
	return result

}
