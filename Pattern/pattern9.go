package main

import "fmt"

func Pattern7(num int) {

	for i := 0; i < num; i++ {

		// Space
		for j := 0; j < num-i-1; j++ {

			fmt.Print(" ")

		}

		// Star
		for j := 0; j < i*2+1; j++ {

			fmt.Print("*")

		}

		// Space
		for j := 0; j < num-i-1; j++ {

			fmt.Print(" ")

		}

		fmt.Println()

	}
}

func Pattern8(num int) {

	for i := 0; i < num; i++ {

		// Space
		for j := 0; j < i; j++ {

			fmt.Print(" ")

		}

		// Star
		for j := 0; j < 2*num-(2*i+1); j++ {

			fmt.Print("*")

		}

		// Space
		for j := 0; j < i; j++ {

			fmt.Print(" ")

		}

		fmt.Println()

	}
}
func main() {

	fmt.Println("Enter the number ")

	var n int
	fmt.Scanln(&n)

	Pattern7(n)
	Pattern8(n)
}
