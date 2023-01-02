package test

import "testing"

func TestIsEven(t *testing.T) {

	expected := false

	result := IsEven(6)

	if expected != result {
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf(" SUCCEDED, expected -> %v, got -> %v", expected, result)

	}

}
