package main

import (
	"fmt"
)

func MinimumInt(first int, rest ...int) int {
	min := first
	for _, x := range rest {
		if x < min {
			min = x
		}
	}
	return min
}

func MinimumFloat(first float64, rest ...float64) float64 {
	min := first
	for _, x := range rest {
		if x < min {
			min = x
		}
	}
	return min
}

func Minimum(first interface{}, rest ...interface{}) interface{} {
	min := first
	for _, x := range rest {
		switch x := x.(type) {
		case int:
			if x < min.(int) {
				min = x
			}
		case float64:
			if x < min.(float64) {
				min = x
			}
		}
	}
	return min
}

func main() {
	fmt.Println("Generics (int)", MinimumInt(10, 22, 3, 5))
	fmt.Println("Generics (float64)", MinimumFloat(10.0, 22.1, 3.2, 5.4))
	fmt.Println("Generics (int interface)", Minimum(10, 22, 3, 5))
	fmt.Println("Generics (float64 interface)", Minimum(10.0, 22.1, 3.2, 5.4))
}
