// Go makes it possible to recover from a panic, by using the recover built-in function.
// A recover can stop a panic from aborting the program and let it continue with execution
// instead.

package main

import "fmt"

func handlePanic() {

	if a := recover(); a != nil {
		fmt.Println("Recover")

	}
}

func entry(fName *string, lName *string) {

	defer handlePanic()
	if fName == nil {
		fmt.Println("First Name can't be nil")
	}
	if lName == nil {
		fmt.Println("Last Name can't be nil")
	}

	fmt.Printf("First Name %s \n Last Name %s \n ", *fName, *lName)
}

func main() {

	fName := "Prashant"
	entry(&fName, nil)

	fmt.Println("The end of main function")

}
