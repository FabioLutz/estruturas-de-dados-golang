package main

import (
	"fmt"

	"github.com/FabioLutz/data-structure/pkg/linkedlist"
)

func main() {
	list := linkedlist.New(1)

	fmt.Println(list)
	
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	
	fmt.Println(list)
}
