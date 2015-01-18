//
// Dynamic function is a function implement the same functionality using
// different algorithms that can be set at runtime in order to a selector
//

package main

import (
	"fmt"
	"os"
)

var myFunc func(string) bool

func main() {
	myFunc = func(s string) bool {
		fmt.Println(s)
		return true
	}
	if os.Args[1] == "1" {
		myFunc = func(s string) bool {
			fmt.Println(s)
			return false
		}
	}
	fmt.Println(myFunc("sample"))
}
