package main

import (
	"fmt"
	"time"
)

/*
The worker pool pattern involves creating a group of worker goroutines to process
tasks concurrently. This pattern is valuable when you have a large number of tasks to execute.
*/

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worked id ", id, "started jobs : ", j)
		time.Sleep(time.Second)
		fmt.Println("worked id ", id, "finished jobs : ", j)
		result <- j * 2

	}
}

func main() {
	const noOfJobs = 5
	job := make(chan int, noOfJobs)
	result := make(chan int, noOfJobs)

	// creates 3 goroutines for the worker function
	for i := 1; i <= 3; i++ {
		go worker(i, job, result)
	}

	// sends numJobs integers to the jobs channel
	for j := 1; j <= noOfJobs; j++ {
		job <- j
	}
	close(job)

	// receives numJobs integers from the results channel
	for k := 1; k <= noOfJobs; k++ {
		<-result
	}
	close(result)

}
