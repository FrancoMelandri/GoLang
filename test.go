package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rnd := rand.Intn(10)
	fmt.Println(rnd)
	rnd = rand.Intn(10)
	fmt.Println(rnd)
}
