package main

import (
	"fmt"
)

var msg = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
	"- Elliot Alderson",
	"Mr. Robot",
}

func Producer(ch chan<- string) {
	for _, v := range msg {
		ch <- v
	}
	close(ch)

}

func Consumer(worker int, ch <-chan string, done chan<- bool) {

	for msg := range ch {
		fmt.Printf("Message %v is consumed by worker %v.\n", msg, worker)
	}
	done <- true
}

// This main goroutine does not wait for any other goroutines to finish.
// But we do not want that — we want our producer-consumer to do their job.
// So, we use the done channel to ensure that the main goroutine can only exit
// when others are done.
func main() {
	c := make(chan string)
	done := make(chan bool)

	go Producer(c)

	for i := 1; i <= 3; i++ {
		go Consumer(i, c, done)

	}

	<-done
}
