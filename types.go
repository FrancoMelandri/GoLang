package main

import (
	"fmt"
)

type coordinate int

type Point struct {
	x coordinate
	y coordinate
}

func (p Point) Show() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p *Point) Show1() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func main() {

	// object createion
	p := Point{1, 2}
	p1 := &Point{1, 2}

	fmt.Println("Hello type", p)

	fmt.Println(p.Show())

	fmt.Println(p1.Show())
	fmt.Println(p1.Show1())
}
