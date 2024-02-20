package leetcode

import (
	"sort"
)

func FindLeastNumOfUniqueInts(arr []int, k int) int {

	// find count of each numbers
	mp := make(map[int]int)
	for _, v := range arr {
		if _, ok := mp[v]; !ok {
			mp[v] = 1
		} else {
			mp[v] += 1
		}
	}
	// sort the keys on the basis of their count
	/*
		Logic should be we area going to remove first k element which have least number of count
	*/

	arrOfObject := []map[int]int{}
	for key, val := range mp {
		arrOfObject = append(arrOfObject, map[int]int{key: val})
	}
	// Sort by least number ok key
	sort.Slice(arrOfObject, func(i, j int) bool {
		a := arrOfObject[i]
		b := arrOfObject[j]
		aValue := 0
		bValue := 0
		for _, v := range a {
			aValue = v
		}
		for _, v := range b {
			bValue = v
		}
		return bValue > aValue
	})

	// Remove the k element from the array

	for _, arrv := range arrOfObject {
		if k == 0 {
			break
		}
		for key, v := range arrv {
			if v <= k {
				k -= v
				arrv[key] = 0
			} else {

				arrv[key] -= k
				k = 0
			}
		}
	}

	count := 0
	for _, v := range arrOfObject {
		value := -1
		for _, v := range v {
			value = v
		}
		if value > 0 {
			count++
		}
	}

	// now filter out with

	return count

}
