package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const JobsLen int = 2
const MaxInteger int = 30
const Verbose bool = true

// Job simple struct with name
// and time to wait
type Job struct {
	Name    string
	Compute int
	done    chan bool
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func startSerialized(jobList []Job) {
	for _, job := range jobList { // Blocks waiting for a send
		if Verbose {
			fmt.Println(job) // Do one job
		}
		job.Compute = fibonacci(job.Compute)
		if Verbose {
			fmt.Println("Ended", job.Compute) // Do one job
		}
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
			job.Compute = fibonacci(job.Compute)
			if Verbose {
				fmt.Println("Ended", job.Compute) // Do one job
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
		jobList[i] = Job{fmt.Sprintf("Job%d", i), rand.Intn(MaxInteger), make(chan bool)}
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
