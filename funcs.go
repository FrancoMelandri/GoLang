package main

import (
	"fmt"
)

//
// variadic
func Minimum(first int, rest ...int) int {
	for _, x := range rest {
		if x < first {
			first = x
		}
	}
	return first
}

//
// Multiple return values
func Swap(first int, second int) (int, int) {
	first, second = second, first
	return first, second
}

func Swap1(first int, second int) (out1 int, out2 int) {
	out1, out2 = second, first
	return
}

func main() {

	// variadic
	fmt.Println(Minimum(10, 2, 5, 7))

	// multiple value
	fmt.Println(Swap(10, 2))
	fmt.Println(Swap1(9, 5))

	// anonymous function or literal
	foo := func() {
		fmt.Println("hello ")
	}
	foo()
}
