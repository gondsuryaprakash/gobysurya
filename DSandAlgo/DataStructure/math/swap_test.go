package math

import "testing"

func TestSwap(t *testing.T) {

	expected1 := 5
	expected2 := 10

	result1, result2 := SwapNumber(10, 5)

	if result1 == expected1 && result2 == expected2 {
		t.Log("Pass")
	} else {
		t.Error("Fail", "Expected1--->", expected1, "result1---> ", result1, "Expected2 -->", expected2, "Result2---->", result2)
	}
}
