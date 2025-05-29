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

	fmt.Println(array.Get(1))

	newArray := dynamicarray.NewWith(1,2,3)

	newArray.Set(0, 2)

	fmt.Println(newArray.Get(0))

	newArray.Set(2, 7)
	
	fmt.Println(newArray.Get(2))

	newArray.AddIn(1,83)

	fmt.Println(newArray)

	newArray.AddIn(3,99)

	fmt.Println(newArray)

	newArray.Remove()
	newArray.Remove()

	fmt.Println(newArray)

	newArray.AddIn(2,142)

	fmt.Println(newArray)

	fmt.Printf("Length: %d\n", newArray.Len())

	newArray2 := dynamicarray.NewWith(1,2,3,4,5,6)

	fmt.Println(newArray2)

	newArray2.RemoveIn(0)

	fmt.Println(newArray2)

	newArray2.RemoveIn(3)

	fmt.Println(newArray2)

	newArray3 := dynamicarray.NewLength(10)

	fmt.Println(newArray3)

	fmt.Printf("Length: %d\n", newArray3.Len())
}
