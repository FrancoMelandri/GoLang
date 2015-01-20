package main

import (
	"fmt"
)

func InsertStringSliceCopy(slice, insertion []string, index int) []string {

	result := make([]string, len(slice)+len(insertion))

	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)

	// built-in function that copy two slides
	copy(result[at:], slice[index:])

	return result

	//return append(slice[:index], append(insertion, slice[index:]...)...)
}

func main() {
	// array
	arr := []string{"A", "B", "C"}

	// initial array
	fmt.Println("Hello append ", arr)

	// built-in append function to append items
	arr = append(arr, "D", "E", "F")
	// array
	fmt.Println("Hello append ", arr)

	// append canc create a new slice if the orinila slice
	// capacity is not enaought to caontain the new item
	arr1 := InsertStringSliceCopy(arr, []string{"x", "y"}, 3)
	fmt.Println("Hello append ", arr1)
}
