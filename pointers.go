package main

import (
	"fmt"
)

func main() {

	var i int

	i = 10
	j := &i
	k := &j

	//i = 11
	//*j++
	//**k++

	//j++ // is not possible to disallow trouble

	fmt.Println("Hello pointers ", i, *j, **k)
}
