//
// A higher order function is a function that takes one or more other functions as
// arguments and uses them in its own body
//

package main

import (
	"fmt"
)

func myFunc(ho func() bool) {
	fmt.Print("called:")
	ho()
}

func main() {
	myFunc(func() bool {
		fmt.Println("func1")
		return true
	})

	myFunc(func() bool {
		fmt.Println("func2")
		return false
	})

	f := func() bool {
		fmt.Println("func3")
		return true
	}
	myFunc(f)
}
