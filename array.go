package main

import (
	"fmt"
)

func main() {
	var a1 [10]int
	a2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a3 []int
	fmt.Println("Hello array ", a1, a2, len(a2), cap(a2), cap(a3))
}
