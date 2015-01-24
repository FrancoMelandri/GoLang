//
// first-class values; is possible to store function
// in a variable
//

package main

import (
	"fmt"
	"time"
)

func myFunc(val int) {
	fmt.Println("myFunc:", val)
}

func main() {
	f := myFunc
	defer func() {
		f(10)
	}()
	go func() {
		f(30)
	}()
	myFunc(20)
	//time.Sleep(1)
}
