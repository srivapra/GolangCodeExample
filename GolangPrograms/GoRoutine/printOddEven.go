package main

import "fmt"

func PrintNumbers(odd, even <-chan int) {

	for {
		select {
		case oddNumber := <-odd:
			fmt.Print(oddNumber, " ")

		case evenNumber := <-even:
			fmt.Print(evenNumber, " ")

		}
	}

}
func main() {

	oddChan := make(chan int)
	eveChan := make(chan int)

	go PrintNumbers(oddChan, eveChan)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			oddChan <- i
		} else {
			eveChan <- i
		}
	}

}
