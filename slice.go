package main

import (
	"fmt"
)

func main() {
	// array
	a1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// slice
	s := a1[1:3]

	fmt.Println("Hello array ", a1, s)

	//	s[0] = 99

	//	fmt.Println("Hello array ", a1, s)
}
