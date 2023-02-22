package singlton

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetInstance(t *testing.T) {

	counter := GetInstance()

	if counter == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	currentCount := counter.AddOne()

	if currentCount != 1 {
		t.Error("After calling for the first time to count, the count must be 1 but ")
	} else {
		assert.Equal(t, currentCount, 1)
	}

	expectedCounter := counter

	counter2 := GetInstance()

	if expectedCounter != counter2 {
		t.Error("Expected same instance in counter2 but it got a different instance")
	}
	currentCounter2 := counter2.AddOne()

	if currentCounter2 != 2 {
		t.Error("After calling for the first time to count, the count must be 2 but it is")
	} else {
		assert.Equal(t, currentCounter2, 2)
	}

}
