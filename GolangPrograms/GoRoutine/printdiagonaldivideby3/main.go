package main

import (
	"fmt"
	"sync"
)

func PassDiagonalToChannel(arr [][]int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j && arr[i][j]%3 == 0 {
				ch <- arr[i][j]

			}
		}
	}
	close(ch)

}

func PrintDiagonal(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range ch {
		fmt.Println(n)
	}

}

func main() {

	arr := [][]int{{9, 2, 3},
		{4, 15, 6},
		{7, 8, 9}}
	fmt.Print(arr)

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go PrintDiagonal(ch, &wg)
	go PassDiagonalToChannel(arr, ch, &wg)
	wg.Wait()

}
