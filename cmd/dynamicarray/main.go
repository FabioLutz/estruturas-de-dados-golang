package main

import (
	"fmt"

	"github.com/FabioLutz/data-structure/pkg/dynamicarray"
)

func main() {
	array := dynamicarray.New()

	fmt.Println(array)

	array.Add(1)

	fmt.Println(array)

	array.Add(2)

	fmt.Println(array)

	array.Add(3)

	fmt.Println(array)
}
