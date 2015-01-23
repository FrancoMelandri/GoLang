package main

import (
	"fmt"
	"strings"
)

// factory function returns a closure which has captured
// the suffix variable at the time the closure was created
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addZip := MakeAddSuffix(".zip")
	addPng := MakeAddSuffix(".png")

	fmt.Println("Hello world ", addZip("sample"), addPng("sample"))
}
