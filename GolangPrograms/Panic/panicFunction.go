// we can create your own panic using panic function

package main

import "fmt"

func getName(name *string, address *string) {

	if name == nil {
		panic("Name cannot be nil")

	}
	if address == nil {
		panic("Address cannot be nil")

	}

	fmt.Printf("Name %s \n Address %s \n", *name, *address)

}

func main() {

	name := "Amit"
	getName(&name, nil)
}
