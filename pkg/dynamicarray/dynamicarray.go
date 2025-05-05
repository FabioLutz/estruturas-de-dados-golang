package dynamicarray

import "fmt"

type dynamicArray struct {
	data []int
}

func New() *dynamicArray {
	return &dynamicArray{
		data: make([]int, 0),
	}
}

func NewWith(values ...int) *dynamicArray {
	return &dynamicArray{
		data: values,
	}
}

func (dynamicArray *dynamicArray) String() string {
	return fmt.Sprintf("%v", dynamicArray.data)
}

func (dynamicArray *dynamicArray) Add(value int) *dynamicArray {
	var length int = len(dynamicArray.data)

	var newArray = make([]int, length+1)

	copy(newArray, dynamicArray.data)

	newArray[length] = value

	dynamicArray.data = newArray

	return dynamicArray
}

func (dynamicArray *dynamicArray) RemoveFirst() *dynamicArray {
	var length int = len(dynamicArray.data)

	if length == 0 {
		panic("failed to remove value: dynamic array is empty")
	}

	var newArray = make([]int, length-1)

	copy(newArray, dynamicArray.data[:length-1])

	dynamicArray.data = newArray

	return dynamicArray
}
