package main

import "fmt"

func sendData(ch chan string, s string) {

	ch <- s
	close(ch)

}

func main() {

	ch := make(chan string)

	str := "Hello"

	go sendData(ch, str)

	value, ok := <-ch
	if !ok {
		fmt.Println("Channel Closed")
	}
	fmt.Println(value)
}
