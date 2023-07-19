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

func Consumer(ch <-chan string, idx int, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
		fmt.Printf("Consumed message \"%v\"\n", msg)
	}

}

func main() {
	c := make(chan string)
	wgP := sync.WaitGroup{}
	wgC := sync.WaitGroup{}

	wgP.Add(4)
	wgC.Add(3)

	for i := 0; i < 4; i++ {
		go Producer(c, i, &wgP)

	}
	for i := 0; i < 3; i++ {
		go Consumer(c, i, &wgC)

	}

	wgP.Wait()
	close(c)
	wgC.Wait()
}
