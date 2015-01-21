package main

import (
	"fmt"
)

func main() {

	// standard for
	for i := 0; i < 5; i++ {
		fmt.Println("Hello for", i)
	}

	// for with break to exit
FOUND:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break FOUND
		}
	}
}
