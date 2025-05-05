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
