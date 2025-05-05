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

	newArray := dynamicarray.NewWith(1,2,3)
	
	newArray.AddIn(1,83)

	fmt.Println(newArray)

	newArray.AddIn(3,99)

	fmt.Println(newArray)

	newArray.RemoveFirst()
	newArray.RemoveFirst()

	fmt.Println(newArray)

	newArray.AddIn(2,142)

	fmt.Println(newArray)
}
