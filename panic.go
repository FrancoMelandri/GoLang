package main

import (
	"fmt"
)

func ConvertByteToInt(x byte) int {
	if 0 <= x && x <= 127 {
		return int(x)
	}
	panic(fmt.Sprintf("%d is out of the range", x))
}

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Error", e)
		}
	}()

	c1 := ConvertByteToInt(10)
	fmt.Println("first conversion", c1)
	c2 := ConvertByteToInt(200)
	fmt.Println("second conversion", c2)
}
