package main

import "fmt"

func FibonacciSeries(n int) {
	n1 := 0
	n2 := 1
	fmt.Print(n1, " ", n2, " ")

	for i := 2; i < n; i++ {
		n3 := n1 + n2
		n1 = n2
		n2 = n3
		fmt.Print(n2, " ")

	}
}

func main() {
	FibonacciSeries(10)

}
