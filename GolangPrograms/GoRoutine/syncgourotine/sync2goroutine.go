package main

import (
	"fmt"
	"sync"
)

func PrintEven(turn chan bool, num chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for n := range num {
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

	for i := 0; i <= 10; i++ {
		num <- i
		<-turn
	}

	close(turn)
	close(num)

	wg.Wait()

}
