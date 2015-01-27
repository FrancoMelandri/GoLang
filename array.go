package main

import (
	"fmt"
)

func main() {
	var a1 [10]int
	a2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a3 []int
	fmt.Println("Hello array ")
	fmt.Println("a1:", a1, len(a1), cap(a1))
	fmt.Println("a2:", a2, len(a2), cap(a2))
	fmt.Println("a3:", a3, cap(a3), cap(a3))
}
