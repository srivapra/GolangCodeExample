package main

import (
	"fmt"
)

func PrintRightAngleTrainglePattern(n int) {

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func main() {

	fmt.Println("Enter the number ")
	var n int
	fmt.Scanln(&n)

	PrintRightAngleTrainglePattern(n)

}
