package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 1; i <= 10; i++ {
		go func(num int) {
			ch <- 2 * num
		}(i)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}
