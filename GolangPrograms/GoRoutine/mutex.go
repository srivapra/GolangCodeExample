// Mutex in golang is used to avoid the race condition in your program
// Two methods defined on Mutex. Lock() and Unlock() to lock the mutex and
// unlock the mutex simultaneously

// Any code present between a call to Lock and Unlock will be executed
// by only one Goroutine at a time

package main

import (
	"fmt"
	"sync"
)

var GFG int = 0

// This is the function we’ll run in every
// goroutine. Note that a WaitGroup must
// be passed to functions by pointer.
func worker(wg *sync.WaitGroup, m *sync.Mutex) {

	// Lock() the mutex to ensure
	// exclusive access to the state,
	// increment the value,
	// Unlock() the mutex
	m.Lock()
	GFG = GFG + 1
	m.Unlock()
	wg.Done()

}

func main() {

	var wg sync.WaitGroup
	var mut sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &mut)
	}
	wg.Wait()
	fmt.Println("Value of x", GFG)

}
