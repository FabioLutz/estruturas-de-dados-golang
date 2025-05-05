package dynamicarray

import "fmt"

type dynamicArray struct {
	data   []int
	length int
}

func New() *dynamicArray {
	return &dynamicArray{
		data:   make([]int, 0),
		length: 0,
	}
}

func NewLength(length uint) *dynamicArray {
	return &dynamicArray{
		data:   make([]int, length),
		length: int(length),
	}
}

func NewWith(values ...int) *dynamicArray {
	return &dynamicArray{
		data:   values,
		length: len(values),
	}
}

func (dynamicArray *dynamicArray) String() string {
	return fmt.Sprintf("%v", dynamicArray.data)
}

func (dynamicArray *dynamicArray) Get(index uint) int {
	if int(index) >= dynamicArray.length {
		panic(fmt.Sprintf("failed to add: index %d out of bounds [0:%d]", index, dynamicArray.length))
	}
	return dynamicArray.data[index]
}

func (dynamicArray *dynamicArray) Set(index uint, value int) {
	if int(index) >= dynamicArray.length {
		panic(fmt.Sprintf("failed to add: index %d out of bounds [0:%d]", index, dynamicArray.length))
	}

	dynamicArray.data[index] = value
}

func (dynamicArray *dynamicArray) Add(value int) {
	var newArray = make([]int, dynamicArray.length+1)

	copy(newArray, dynamicArray.data)

	newArray[dynamicArray.length] = value

	dynamicArray.length++
	dynamicArray.data = newArray
}

func (dynamicArray *dynamicArray) AddIn(index uint, value int) {
	if int(index) >= dynamicArray.length {
		panic(fmt.Sprintf("failed to add: index %d out of bounds [0:%d]", index, dynamicArray.length))
	}

	var newArray1 = make([]int, index+1)
	var newArray2 = make([]int, dynamicArray.length-int(index))

	copy(newArray1, dynamicArray.data[:index])
	copy(newArray2, dynamicArray.data[index:])

	dynamicArray.data = append(newArray1, newArray2...)

	dynamicArray.length++
	dynamicArray.data[index] = value
}

func (dynamicArray *dynamicArray) RemoveFirst() {
	if dynamicArray.length == 0 {
		panic("failed to remove value: dynamic array is empty")
	}

	var newArray = make([]int, dynamicArray.length-1)

	copy(newArray, dynamicArray.data[:dynamicArray.length-1])

	dynamicArray.length--
	dynamicArray.data = newArray
}

func (dynamicArray *dynamicArray) RemoveIn(index uint) {
	if int(index) >= dynamicArray.length {
		panic(fmt.Sprintf("failed to add: index %d out of bounds [0:%d]", index, dynamicArray.length))
	}

	var newArray1 = make([]int, index)
	var newArray2 = make([]int, dynamicArray.length-int(index)-1)

	copy(newArray1, dynamicArray.data[:index])
	copy(newArray2, dynamicArray.data[index+1:])

	dynamicArray.length--
	dynamicArray.data = append(newArray1, newArray2...)
}
