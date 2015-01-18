package main

import (
	"fmt"
)

func main() {

	i := 10
	s := "World"
	b := true
	c := false
	f := 10.27

	fmt.Printf("Hello Integer %T->%d %b\n", i, i, i)
	fmt.Printf("Hello %T->%d\n", i, i)

	fmt.Printf("Hello %s %c\n", s, s[0])

	fmt.Printf("Hello Boolean %t %t\n", b, c)

	fmt.Printf("Hello Float %f %5.1f %010.05f\n", f, f, f)
}
