package main

import (
	"fmt"
)

func swap(i, j int) (int, int) {
	return j, i
}

func main() {

	i := 10
	j := 20

	fmt.Println("Hello pointers : ", i, j)

	h, k := swap(i, j)

	fmt.Println("Hello pointers swap : ", h, k)
}
