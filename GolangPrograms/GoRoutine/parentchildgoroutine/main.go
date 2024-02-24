package main

import (
	"fmt"
	"sync"
)

func Parent(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	side := 3
	areaOfSquare := side * side
	ch <- areaOfSquare
	close(ch)

}

func Child(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Println("Area is : ", n)
	}

}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go Parent(ch, &wg)
	go Child(ch, &wg)

	wg.Wait()

}
