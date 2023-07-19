// Concurrency is about dealing with lots of things at once.
// Parallelism is about doing lots of things at once.

// Why Goroutine?

// To achieve concurrency and parallelism in Golang,
// we can use the concept of Goroutines.

// Suppose we have a program that has two independent functions. In normal programs,
// two functions will be executed synchronously. That is, the second function is executed
// only after the execution of the first function. However, since both functions are
// independent, it would be efficient if we could execute both functions together
// asynchronously. For this we can use goroutine

// A goroutine is a lightweight thread managed by the Go runtime

package main

import (
	"fmt"
	"time"
)

func print(s string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(s)
	}
}

func main() {

	go print("Hello")
	go print("World")

	// To sleep main goroutine 2 second. Becoz if main goroutine
	// terminate, the above 2 newly created gorutine will never execute
	time.Sleep(2 * time.Second)

}
