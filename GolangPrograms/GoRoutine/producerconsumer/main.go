package main

import (
	"fmt"
	"sync"
)

type Info struct {
	FirstName string
	LastName  string
}

func producer(ch chan Info, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 1; i <= 10; i++ {
		info := Info{
			FirstName: fmt.Sprintf("First %d ", i),
			LastName:  fmt.Sprintf("Last %d ", i),
		}
		ch <- info
	}
	close(ch)

}

func consumer(ch chan Info, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Println(n)
	}

}

func main() {

	var wg sync.WaitGroup
	ch := make(chan Info)

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()

}
