package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Init world!")
}

func main() {
	fmt.Println("Hello world! ", time.Now())
}
