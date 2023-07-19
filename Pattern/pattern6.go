package main

import (
	"fmt"
)

func Pattern(num int) {

	for i := 0; i < num; i++ {
		for j := 1; j <= num-i; j++ {

			fmt.Print(j, " ")

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
