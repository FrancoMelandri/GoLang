package main

import (
	"fmt"
)

func main() {
	// array
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0

	// a is a copy of the item in the slice
	for _, a := range arr {
		sum += a
	}
	fmt.Println("Hello iteration ", sum)

	// iterate a portion of the initial array
	sum = 0
	for _, a := range arr[:5] {
		sum += a
	}
	fmt.Println("Hello iteration ", sum)
}
