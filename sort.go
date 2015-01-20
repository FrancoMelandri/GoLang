package main

import (
	"fmt"
	"sort"
)

func main() {
	// array
	arr := []string{"B", "C", "A"}

	// initial array
	fmt.Println("Hello sort ", arr)

	// sort function standard library
	sort.Strings(arr)
	fmt.Println("Hello sort ", arr)

	// sort function uses the binary comparison log2(n) instead of
	// linear
	// we have to pass the lenght of the slice and a function (clojure) that compare an
	// item with the tagert
	target := "B"
	compare := func(i int) bool { return arr[i] >= target }
	s := sort.Search(len(arr), compare)
	fmt.Println("Hello sort ", s, arr[s])
}
