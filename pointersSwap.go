package main

import (
	"fmt"
)

func swap(i, j *int) {
	*i, *j = *j, *i
}

func main() {

	i := 10
	j := 20

	fmt.Println("Hello pointers : ", i, j)

	swap(&i, &j)

	fmt.Println("Hello pointers swap : ", i, j)
}
