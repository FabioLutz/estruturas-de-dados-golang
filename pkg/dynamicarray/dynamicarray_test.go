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

func TestGet(t *testing.T) {
	var tests = []struct {
		name        string
		array       dynamicArray
		input       int
		expected    int
		expectPanic bool
		panicMsg    string
	}{
		{"valid array", dynamicArray{[]int{4, 2, 6, 3, 7}, 5}, 2, 6, false, ""},
		{"empty array", dynamicArray{make([]int, 0), 0}, 0, -1, true, "failed to get value: index 0 out of bounds [0:0]"},
		{"out of range", dynamicArray{[]int{1, 2, 3}, 3}, 10, -1, true, "failed to get value: index 10 out of bounds [0:3]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recover := recover(); recover != nil {
					if !tt.expectPanic {
						t.Errorf("Unexpected Panic: %v", recover)
					}
					if message, ok := recover.(string); ok {
						if message != tt.panicMsg {
							t.Errorf("Incorrect panic message: Expected %s, Actual %s", tt.panicMsg, message)
						}
					} else {
						t.Errorf("Panic is not a string: %v", recover)
					}
				} else {
					if tt.expectPanic {
						t.Errorf("Expected error, none received")
					}
				}
			}()

			actual := tt.array.Get(uint(tt.input))

			if actual != tt.expected {
				t.Errorf("Incorrect value: Expected %d, Actual %d", tt.expected, actual)
			}
		})
	}
}
