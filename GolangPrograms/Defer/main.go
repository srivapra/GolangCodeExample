// In Golang, the defer keyword is used to delay the execution of a function or a
// statement until the nearby function returns. In simple words, defer will move the
// execution of the statement to the very end inside a function.

package main

import "fmt"

func main() {

	defer fmt.Println("Hello")
	defer fmt.Println("World")

	myDefer()

	fmt.Println("Defer")
}

func myDefer() {
	for i := 0; i < 5; i++ {

		defer fmt.Print(i, " ")

	}
}
