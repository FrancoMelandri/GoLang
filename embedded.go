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
	Item       // embedding
	item *Item // aggregation
	id   int   // aggeegation
}

func (item *Item) Cost() int {
	return item.Price * item.qty
}

func main() {

	// object createion
	i := SpecialItem{Item{"item1", 3, 5}, &Item{"item2", 3, 3}, 1}

	fmt.Println("Hello embedding and aggregating", i)

	fmt.Println("Hello embedding", i.Cost())
	fmt.Println("Hello aggregating", i.item.Cost())

	fmt.Println("Hello embedding", i.Item.Name)
	fmt.Println("Hello aggregating", i.item.Name)

}
