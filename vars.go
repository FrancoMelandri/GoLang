package main

import (
	"fmt"
)

// CONSTs....
//const msg string = "hello"
const msg = "hello"

//
// Enumerations rather then repeat the const string
const (
	Zero = 0
	One  = 1
	Two  = 2
)

func main() {

	// var msg string
	// msg = "hello"

	//msg := "hello"

	var i rune
	var j int32

	i = 12
	j = 24

	fmt.Println(msg, Zero, i, j)
}
