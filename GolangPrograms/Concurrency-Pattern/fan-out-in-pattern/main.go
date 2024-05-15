package main

import (
	"fmt"
	"sync"
)

/*
The fan-out/fan-in pattern involves distributing tasks to multiple worker
goroutines (fan-out) and then aggregating their results (fan-in).
It’s useful for parallelizing tasks and combining their outcomes.
*/

func main() {
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))
	for _, num := range data {
		input <- num
	}
	close(input)

	results := make(chan int, len(data))

	var wg sync.WaitGroup

	// Fan-out: Launch multiple worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range input {
				ans := d * 2
				results <- ans
			}

		}()
	}

	// Fan-in: Aggregate results from workers
	go func() {
		wg.Wait()
		close(results)

	}()

	// Process aggregated results
	for result := range results {
		fmt.Print(result, " ")
	}
}
