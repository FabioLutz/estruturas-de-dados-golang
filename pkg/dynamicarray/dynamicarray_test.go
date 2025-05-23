package dynamicarray

import (
	"testing"
)

func TestNew(t *testing.T) {
	actualArray := New()
	
	expectedLength := 0

	if actualArray.length != expectedLength {
		t.Errorf("Incorrect length field: Expected %d, Actual %d", expectedLength, actualArray.length)
	}

	if len(actualArray.data) != expectedLength {
		t.Errorf("Incorrect data length: Expected %v, Actual %v", expectedLength, len(actualArray.data))
	}
}
