package main

import (
	"fmt"
)

func main() {

	a1 := make([]int, 10, 20)
	//a1 := []int{}
	fmt.Println("Hello array 1", a1, len(a1), cap(a1))
}
