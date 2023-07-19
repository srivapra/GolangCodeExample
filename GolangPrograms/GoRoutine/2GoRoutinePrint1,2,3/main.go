package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println("Even ", i)
			}(i)
			wg.Wait()
		} else {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println("Odd ", i)
			}(i)
			wg.Wait()
		}
	}
}
