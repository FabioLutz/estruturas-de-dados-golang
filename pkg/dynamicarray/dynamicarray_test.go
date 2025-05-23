package dynamicarray

import (
	"reflect"
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

func TestNewLength(t *testing.T) {
	expectedLength := 6
	expectedData := []int{0, 0, 0, 0, 0, 0}

	actualArray := NewLength(uint(expectedLength))

	if len(actualArray.data) != expectedLength {
		t.Errorf("Incorrect array length: Expected %d, Actual %d", expectedLength, len(actualArray.data))
	}

	if actualArray.length != expectedLength {
		t.Errorf("Incorrect length field: Expected %d, Actual %d", expectedLength, actualArray.length)
	}

	if !reflect.DeepEqual(actualArray.data, expectedData) {
		t.Errorf("Incorrect data: Expected %v, Actual %v", expectedData, actualArray.data)
	}
}

func TestNewWith(t *testing.T) {
	expectedData := []int{1, 2, 3, 4, 5, 6}
	expectedLength := len(expectedData)

	actualArray := NewWith(expectedData...)

	if !reflect.DeepEqual(actualArray.data, expectedData) {
		t.Errorf("Incorrect data: Expected %v, Actual %v", expectedData, actualArray.data)
	}

	if len(actualArray.data) != expectedLength {
		t.Errorf("Incorrect array length: Expected %d, Actual %d", expectedLength, len(actualArray.data))
	}

	if actualArray.length != expectedLength {
		t.Errorf("Incorrect length field: Expected %d, Actual %d", expectedLength, actualArray.length)
	}
}
