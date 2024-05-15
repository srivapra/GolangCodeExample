package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker: ", id, "started job: ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker: ", id, "finished job: ", j)
		result <- j * 2
	}
}

func main() {
	startTime := time.Now()
	const numberOfJobs = 5
	job := make(chan int, numberOfJobs)
	result := make(chan int, numberOfJobs)

	// creates 3 goroutines for the worker function
	for w := 1; w <= 3; w++ {
		go Worker(w, job, result)
	}

	// sends numJobs integers to the jobs channel
	for j := 1; j < numberOfJobs; j++ {
		job <- j
	}
	close(job)

	// receives numJobs integers from the results channel
	for k := 1; k < numberOfJobs; k++ {
		<-result
	}
	close(result)
	fmt.Println("Total time ", time.Since(startTime))

}
