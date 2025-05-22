package linkedlist

import "fmt"

type linkedList struct {
	data int
	next *linkedList
}

func New(value int) *linkedList {
	return &linkedList{
		data: value,
		next: nil,
	}
}

func (linkedList *linkedList) String() string {
	return fmt.Sprintf("%v", linkedList.data)
}
