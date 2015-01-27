package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const JobsLen int = 100000
const MaxTime int = 5
const Verbose bool = false

// Job simple struct with name
// and time to wait
type Job struct {
	Name string
	Wait time.Duration
	done chan bool
}

func startSerialized(jobList []Job) {
	for _, job := range jobList { // Blocks waiting for a send
		if Verbose {
			fmt.Println(job) // Do one job
		}
		time.Sleep(job.Wait * time.Second)
	}
}

func waitSerialized() {

}

func startConcurrent(jobList []Job) {
	for _, job := range jobList { // Blocks waiting for a send
		if Verbose {
			fmt.Println("Go", job)
		}
		go func(job Job) {
			if Verbose {
				fmt.Println("Executed", job) // Do one job
			}
			time.Sleep(job.Wait * time.Second)
			if Verbose {
				fmt.Println("Ended", job) // Do one job
			}
			job.done <- true
		}(job)
	}
}

func waitConcurrent(jobList []Job) {
	for _, job := range jobList { // Blocks waiting for a send
		<-job.done
		if Verbose {
			fmt.Println("Waited", job)
		}
	}
}

func main() {

	start := time.Now()
	var elapsed time.Duration

	jobList := make([]Job, JobsLen)

	for i := 0; i < JobsLen; i++ {
		jobList[i] = Job{fmt.Sprintf("Job%d", i), time.Duration(rand.Intn(MaxTime)), make(chan bool)}
	}

	fmt.Println(os.Args)
	if os.Args[1] == "0" {
		// serialized
		start = time.Now()
		fmt.Println("Jobs serialize stared: ", start)
		startSerialized(jobList)
		waitSerialized()
		elapsed = time.Now().Sub(start)
		fmt.Println("Jobs serialize finished: ", elapsed)
	} else {
		// concurrent
		start = time.Now()
		fmt.Println("Jobs concurrent stared: ", start)
		startConcurrent(jobList)
		waitConcurrent(jobList)
		elapsed = time.Now().Sub(start)
		fmt.Println("Jobs concurrent finished: ", elapsed)
	}
}
