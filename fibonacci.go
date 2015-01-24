package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Print(" ", fibonacci(i))
	}
}
