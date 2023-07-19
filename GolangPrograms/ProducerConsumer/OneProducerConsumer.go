package main

import (
	"fmt"
)

var msg = []string{
	"Hi",
	"Hello",
	"How are you",
}

func Producer(ch chan<- string) {
	for _, v := range msg {
		fmt.Println("Producer ", v)
		ch <- v
	}
	close(ch)

}

func Consumer(ch <-chan string, done chan<- bool) {

	for val := range ch {
		fmt.Println("Consumer ", val)
	}
	done <- true
}

// This main goroutine does not wait for any other goroutines to finish.
// But we do not want that — we want our producer-consumer to do their job.
// So, we use the done channel to ensure that the main goroutine can only exit
// when others are done.
func main() {
	c := make(chan string)
	d := make(chan bool)

	go Producer(c)
	go Consumer(c, d)

	<-d
}
