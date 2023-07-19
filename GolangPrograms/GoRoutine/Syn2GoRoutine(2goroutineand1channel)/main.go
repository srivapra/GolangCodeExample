package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// Goroutine 1: send the numbers 1 to 10 on the channel
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	// Goroutine 2: receive the numbers from the channel and print them
	go func() {
		for num := range ch {
			fmt.Println(num)
		}
	}()

	// Wait for both goroutines to complete
	fmt.Scanln()
}
