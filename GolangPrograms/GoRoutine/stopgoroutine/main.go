package main

import (
	"fmt"
	"time"
)

/*
When the program starts, the myProcess goroutine starts running.
It prints "My Goroutine is running :(" every second because no stop signal has been received.
After 3 seconds, the main goroutine sends a stop signal on the stopChannel.
The myProcess goroutine receives the stop signal, prints "Thanks for stopping my goroutine :)", and exits gracefully.
After a total of 4 seconds, the main goroutine prints "Main Goroutine exited" and terminates.
*/

func myProcess(stopchannel chan bool) {
	for {
		select {
		case <-stopchannel:
			fmt.Println("Thanks for stopping my goroutine")
			return
		default:
			fmt.Println("My goroutine is running")
			time.Sleep(time.Second)
		}

	}

}
func main() {
	ch := make(chan bool)
	go myProcess(ch)
	time.Sleep(3 * time.Second)
	ch <- true
	time.Sleep(time.Second)
	fmt.Println("Main Goroutine exited")

}
