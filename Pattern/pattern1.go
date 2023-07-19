package main

import "fmt"

func PrintSqaureStarPattern(num int) {

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func main() {

	fmt.Println("Enter the number ")
	var n int
	fmt.Scanln(&n)

	PrintSqaureStarPattern(n)

}
