package main

import (
	"fmt"
)

func Pattern(num int) {

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

	Pattern(n)
}
