// The WaitGroup type of sync package, is used to wait for the
// program to finish all goroutine. It uses a counter that specifies the number of goroutines,
// and Wait blocks the execution of the program until the WaitGroup counter is zero.

// The Add method is used to add a counter to the WaitGroup.

// The Done method of WaitGroup is scheduled using a defer statement to decrement the WaitGroup
// counter.

// The Wait method of the WaitGroup type waits for the program to finish all goroutines.

// The Wait method is called inside the main function, which blocks execution until the
// WaitGroup counter reaches the value of zero and ensures that all goroutines are executed.

package main

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup) {

	// This decreases counter by 1
	defer wg.Done()

	fmt.Println("I am inside runner1")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am inside runner2")

}

func execute() {

	wg := new(sync.WaitGroup)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	wg.Add(2)
	go runner1(wg)
	go runner2(wg)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}

func main() {
	execute()
}
