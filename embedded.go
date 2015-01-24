package main

import (
	"fmt"
)

type Item struct {
	Name  string // aggeegation
	Price int    // aggeegation
	qty   int    // aggeegation
}

type SpecialItem struct {
	Item     // embedding
	id   int // aggeegation
}

func (item *Item) Cost() int {
	return item.Price * item.qty
}

func main() {

	// object createion
	i := SpecialItem{Item{"item1", 3, 5}, 1}

	fmt.Println("Hello embedding", i)

	fmt.Println("Hello embedding", i.Cost())
}
