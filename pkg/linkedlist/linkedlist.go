package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func New(value int) *LinkedList {
	newNode := Node{data: value, next: nil}

	return &LinkedList{head: &newNode}
}

func (linkedList *LinkedList) String() string {
	current := linkedList.head
	
	if current == nil {
		return "[]"
	}
	
	var stringList string
	for current != nil {
		stringList += fmt.Sprint(current.data)
		
		if current.next != nil {
			stringList += " -> "
		}
		
		current = current.next
	}
	
	return fmt.Sprint("[" + stringList + "]")
}

func (linkedList *LinkedList) Add(value int) {
	newNode := &Node{data: value, next: nil}

	if linkedList.head == nil {
		linkedList.head = newNode
		return
	}

	current := linkedList.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}
