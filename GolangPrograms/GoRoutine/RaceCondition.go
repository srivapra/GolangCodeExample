// A race condition occurs when multiple goroutines try to access and modify
// the same data/variable (memory address). E.g., if one goroutine tries to increase an
// integer and another goroutine tries to read it, this will cause a race condition.

// Program with race condition
package main

import (
	"fmt"
	"sync" // to import sync later on
)

var GFG int = 0

// This is the function we’ll run in every
// goroutine. Note that a WaitGroup must
// be passed to functions by pointer.
func worker(wg *sync.WaitGroup) {
	GFG = GFG + 1

	// On return, notify the
	// WaitGroup that we’re done.
	wg.Done()
}
func main() {

	// This WaitGroup is used to wait for
	// all the goroutines launched here to finish.
	var w sync.WaitGroup

	// Launch several goroutines and increment
	// the WaitGroup counter for each
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w)
	}
	// Block until the WaitGroup counter
	// goes back to 0; all the workers
	// notified they’re done.
	w.Wait()
	fmt.Println("Value of x", GFG)
}

// Explanation: In the program above, the worker function in line no. 12 increments
// the value of GFG by 1 and then calls Done() on the WaitGroup to inform its completion.
// The worker function is called 1000 times. Each of these Goroutines run simultaneously
// and race condition occurs when trying to increment GFG in line no. 13 as multiple
// Goroutines try to access the value of GFG concurrently. Running the same program
// multiple times gives different outputs each time because of the race condition.
