package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Hello world for strings"

	sli := strings.LastIndex(s, " ")
	si := strings.Index(s, " ")

	firstWord := s[:si]
	lastWord := s[sli+1:]
	cnt := strings.Count(s, " ")
	replaced := strings.Replace(s, " ", "-", cnt)

	fmt.Println(s, sli, si)
	fmt.Println(firstWord)
	fmt.Println(lastWord)
	fmt.Println(replaced)
}
