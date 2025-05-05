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
