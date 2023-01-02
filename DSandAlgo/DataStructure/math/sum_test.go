package main

import "testing"

func TestSum(t *testing.T) {
	expected := 15

	result := SumOfNumber(5)

	if expected != result {
		t.Error("Fail")
	} else {
		t.Log("Pass")
	}

}
