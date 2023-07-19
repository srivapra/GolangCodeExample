package main

import (
	"fmt"
)

func PrintRightAngleTraingleNumberPattern(num int) {

	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {

			fmt.Print(j, " ")

		}
		fmt.Println()
	}
}

func main() {

	fmt.Println("Enter the number ")

	var n int
	fmt.Scanln(&n)

	PrintRightAngleTraingleNumberPattern(n)
}
