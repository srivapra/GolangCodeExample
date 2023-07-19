package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan string)

	go func() {

		defer wg.Done()

		sendMsg := "hello"
		fmt.Println("Meesage sent : ", sendMsg)

		ch <- sendMsg

	}()

	go func() {

		defer wg.Done()
		time.Sleep(2 * time.Second)

		rmsg := <-ch
		fmt.Println("Message recieved :", rmsg)

	}()

	wg.Wait()

}
