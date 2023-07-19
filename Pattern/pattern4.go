package main

import (
	"fmt"
)

func Pattern(num int) {

	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {

			fmt.Print(i, " ")

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
