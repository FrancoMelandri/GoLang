package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var input string

	if os.Args[1] == "0" {

		fmt.Println("Press any key and RETURN")
		input, _ = reader.ReadString('\n')

	} else {

		read := make(chan bool)
		go func() {
			fmt.Println("go routine started")
			fmt.Println("Press any key and RETURN")
			input, _ = reader.ReadString('\n')
			read <- true
		}()
		<-read
	}

	fmt.Println("Hello Concurrent", input)
}
