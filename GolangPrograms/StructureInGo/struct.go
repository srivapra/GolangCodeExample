// A structure or struct in Golang is a user-defined type that allows
// to group/combine items of possibly different types into a single type.
// Any real-world entity which has some set of properties/fields can be represented as a
// struct. This concept is generally compared with the classes in object-oriented programming.

// For Example, an address has a name, street, city, state, Pincode. It makes
// sense to group these three properties into a single structure address.

// Golang program to show how to
// declare and define the struct

package main

import "fmt"

// Defining a struct type
type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a Address
	fmt.Println(a)

	// Declaring and initializing a
	// struct using a struct literal
	a1 := Address{"Prashant", "Panipat", 3623572}

	fmt.Println("Address1: ", a1)

	// Naming fields while
	// initializing a struct
	a2 := Address{Name: "Prashant", city: "Panipat", Pincode: 277001}

	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := Address{Name: "Prashant"}
	fmt.Println("Address3: ", a3)

	// Accessing field of struct using dot operator
	fmt.Println("Accessing field of struct using dot operator: ", a1.Name)
	fmt.Println("Accessing field of struct using dot operator: ", a1.Pincode)

	// address is a pointer to the Employee struct
	address := &Address{"Srivastava", "Haryana", 132103}
	fmt.Println("Pointer to struct : ", (*address).Name)

	// Array of struct
	var arrayStruct = []Address{
		{
			Name:    "Amit",
			city:    "Delhi",
			Pincode: 123,
		},
		{
			Name:    "Sumit",
			city:    "Old Delhi",
			Pincode: 12345,
		},
	}
	fmt.Println("Array Of Struct : ", arrayStruct)

	// Iterating through the array struct and printing the name
	for _, arrayStructValue := range arrayStruct {
		fmt.Println(arrayStructValue.Name)
	}

}
