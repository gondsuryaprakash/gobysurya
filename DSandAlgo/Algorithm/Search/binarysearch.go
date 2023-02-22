package search

import (
	"errors"
	"sort"
)

func BinarySearch(key, low, high int, arr []int) (index int, err error) {

	// Sort the array
	sort.Ints(arr)
	if low > high {
		return -1, errors.New("Key is not Present in Arr")
	}
	mid := low + (high-low)/2
	if arr[mid] == key {
		return mid, nil
	}
	if arr[mid] > key { // Element present in right side
		return BinarySearch(key, low, mid-1, arr)
	} else {
		return BinarySearch(key, mid+1, high, arr)
	}

}
