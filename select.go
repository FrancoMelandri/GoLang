package main

import (
	"fmt"
	"time"
)

func main() {

	// channels creations
	timeChan := time.NewTimer(time.Second).C
	timeTick := time.NewTicker(time.Millisecond * 400).C
	doneChan := make(chan bool)

	// go routine creation
	go func() {
		// start concurrent routine
		fmt.Println("Start wait")
		time.Sleep(time.Second * 2)
		doneChan <- true
	}()

	//
	// Infite loop to wait all the events
	// on channels
	fmt.Println("Loop...")
	for {
		select {
		case <-timeChan:
			fmt.Println("Timer expired")
		case <-timeTick:
			fmt.Println("Time tick")
		case <-doneChan:
			fmt.Println("Done!")
			return
		}
	}
}
