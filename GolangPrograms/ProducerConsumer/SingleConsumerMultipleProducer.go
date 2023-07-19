package main

import (
	"fmt"
	"sync"
)

var messages = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

func Producer(ch chan<- string, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, msg := range messages[idx] {
		fmt.Printf("Producer %v sending message \"%v\"\n", idx, msg)
		ch <- msg
	}

}

func Consumer(ch <-chan string, done chan<- bool) {

	for msg := range ch {
		fmt.Printf("Consumed message \"%v\"\n", msg)
	}
	done <- true
}

// This main goroutine does not wait for any other goroutines to finish.
// But we do not want that — we want our producer-consumer to do their job.
// So, we use the done channel to ensure that the main goroutine can only exit
// when others are done.
func main() {
	c := make(chan string)
	wg := sync.WaitGroup{}
	done := make(chan bool)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go Producer(c, i, &wg)

	}

	go Consumer(c, done)
	wg.Wait()
	close(c)
	<-done
}
