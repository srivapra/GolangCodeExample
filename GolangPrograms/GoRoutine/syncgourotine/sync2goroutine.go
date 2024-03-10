package main

import (
	"fmt"
	"sync"
)

// Receives two channels as arguments: turn for synchronization and num for receiving numbers.
func PrintEven(turn chan bool, num chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for n := range num {

		// If it's even, prints the number and then sends a signal (true) to the turn channel,
		// indicating that it's done printing
		if n%2 == 0 {
			fmt.Println("Even : ", n)
			turn <- true
		} else {
			num <- n
		}
	}

}

func PrintOdd(turn chan bool, num chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range num {
		if n%2 != 0 {
			fmt.Println("Odd : ", n)
			turn <- true
		} else {
			num <- n
		}
	}

}

func main() {

	turn := make(chan bool)
	num := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go PrintEven(turn, num, &wg)
	go PrintOdd(turn, num, &wg)

	// Loops from 0 to 10, sending each number to the num channel and then waiting for a signal
	// on the turn channel (<-turn), which ensures that the even and odd printing alternates.
	for i := 0; i <= 10; i++ {
		num <- i
		<-turn
	}

	close(turn)
	close(num)

	wg.Wait()

}
