package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	map1["hello"] = 10

	fmt.Println("Hello maps ", map1, map1["hello"])

	// retries the value and the existence flag
	// lookup a value, if thekey does not exist return
	// 0 so the flag could be use to check
	v, found := map1["world"]
	fmt.Println("Hello maps ", v, found)
	v1, found1 := map1["hello"]
	fmt.Println("Hello maps ", v1, found1)

	// sfae delete of a key value
	delete(map1, "hello")
	delete(map1, "hello")
	fmt.Println("Hello maps ", map1)

	// iteration
	map1["hello"] = 10
	map1["world"] = 20
	for key, val := range map1 {
		fmt.Println(key, val)
	}
	for key := range map1 {
		fmt.Println(key, map1[key])
	}
}
