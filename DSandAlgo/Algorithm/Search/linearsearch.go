package search

import "errors"

func LinearSearch(key int, array []int) (index int, err error) {

	for i := 0; i < len(array); i++ {
		if array[i] == key {
			index = i
			return index, nil
		}
	}
	return -1, errors.New("key is not present in array")
}
