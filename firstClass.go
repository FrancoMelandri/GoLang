//
// first-class values; is possible to store function
// in a variable
//

package main

import (
	"fmt"
)

func myFunc(val int) {
	fmt.Println("myFunc:", val)
}

func main() {
	f := myFunc
	defer func() {
		f(10)
	}()
	myFunc(20)
}
