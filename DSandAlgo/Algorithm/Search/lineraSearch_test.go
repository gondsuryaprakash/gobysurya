package search

import "testing"

func TestLinearSearch(t *testing.T) {

	expected1 := 2
	var expected2 error
	arr := []int{19, 10, 5, 6}

	result, err := LinearSearch(5, arr)

	if expected1 == result && expected2 == err {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

}
