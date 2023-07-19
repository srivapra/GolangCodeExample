package main

import (
	"fmt"
)

func sendEven(ch chan<- int) {
	for i := 0; i <= 10; i += 2 {
		ch <- i
	}
	close(ch)
}

func sendOdd(ch chan<- int) {
	for i := 1; i <= 10; i += 2 {
		ch <- i
	}
	close(ch)
}

func receive(ch <-chan int) {
	for num := range ch {
		fmt.Printf("%d ", num)
	}
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go sendEven(evenCh)
	go sendOdd(oddCh)

	fmt.Println("Even Numbers:")
	receive(evenCh)

	fmt.Println("\nOdd Numbers:")
	receive(oddCh)
}
