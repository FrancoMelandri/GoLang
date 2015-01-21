package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello procedural", i)
	}

	// short variable declaration
	var a int
	a = 10
	b := 10
	fmt.Printf("%T%d %T%d\n", a, a, b, b)
	// ------------

	// multiple variable definition
	c, d, e, f := 6, 7, 8, 9
	fmt.Println(c, d, e, f)

}
