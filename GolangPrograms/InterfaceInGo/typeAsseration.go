// Type assertions in Golang provide access to the exact type of variable of an interface.
// A type assertion takes an interface value and extracts from it a value of the
// specified explicit type

package main

import "fmt"

func main() {

	// create an empty interface
	var a interface{}

	// store integer to an empty interface
	a = 10

	// type assertion
	interfaceIntValue := a.(int)
	fmt.Println(interfaceIntValue)

	// Will throw panic error
	// interfaceStringValue := a.(string)
	// fmt.Println(interfaceStringValue)

	// Avoiding Panic

	avoidPanic, ok := a.(string)
	fmt.Println(avoidPanic)
	fmt.Println(ok)

}
