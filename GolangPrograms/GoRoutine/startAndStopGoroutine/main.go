package main

import (
	"fmt"
	"time"
)

func sendMsg(ch chan<- string) {
	ticker := time.NewTicker(2 * time.Second)
	defer close(ch)
	defer ticker.Stop()
	stopTimer := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			ch <- "Hi"

		case <-stopTimer.C:
			fmt.Println("Stopping go routine")
			return

		}
	}

}

func main() {
	ch := make(chan string)
	go sendMsg(ch)
	for msg := range ch {
		fmt.Println("Recieved Msg : ", msg)
	}
}
