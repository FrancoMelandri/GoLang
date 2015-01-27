package main

import (
	"fmt"
)

func main() {

	a1 := make([]int, 10, 20)
	//a1 := []int{}
	s := fmt.Sprintf("%T", a1)
	fmt.Println("Hello array 1", s, a1, len(a1), cap(a1))
}
