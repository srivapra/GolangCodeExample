package main

import "fmt"

/*
The pipeline pattern in Go is a concurrency pattern used for efficiently
processing large amounts of data. It involves dividing a complex task into smaller stages,
each executed concurrently by separate goroutines.The output of one stage is passed as input
to the next stage through channels, forming a pipeline.
*/

func main() {
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))

	for _, i := range data {
		input <- i
	}
	close(input)

	// First stage of the pipeline: Doubles the input values
	double := make(chan int)
	go func() {
		defer close(double)
		for num := range input {
			double <- num * 2
		}

	}()

	// Second stage of the pipeline: Squares the doubled values
	square := make(chan int)
	go func() {
		defer close(square)
		for d := range double {
			square <- d * d
		}

	}()

	// Third stage of the pipeline: Prints the squared values
	for v := range square {
		fmt.Print(v, " ")
	}
}
