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

	newArray := dynamicarray.NewWith(3,2,1)

	newArray.Add(9)
	newArray.Add(8)
	newArray.Add(7)

	fmt.Println(newArray)

	newArray.RemoveFirst()
	newArray.RemoveFirst()
	newArray.RemoveFirst()

	fmt.Println(newArray)

	newArray.RemoveFirst()
	newArray.RemoveFirst()
	newArray.Add(5)
	newArray.Add(6)

	fmt.Println(newArray)
}
