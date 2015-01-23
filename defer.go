package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello deferred 1!")
	defer fmt.Println("Hello deferred 2!")
	fmt.Println("Hello world!")
}
