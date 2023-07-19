// A pointer is just a direct reference to a memory address.

// There is a problem in the programming language. Whenever you declare a variable, constant and array
// they are just reference to the memory. These all value that you put inside these variables and constant
// they got stored in some memory and the problem is when sometime you passed these variable in variety of
// functions, sometimes there is a problem that these variable is not actually passed along, a copy of
// being created for these variables and copy is passed into the multiple of functions, It's create
// iragularity in the program. To avoid these iragularity we have the pointers.

// Pointers give you the assuraty that whatever resources you are passing as a pointer
// it is not a copy of the variables, It is the actual value of the memory that is being passed on.

package main

import "fmt"

func main() {

	// Declare a pointer that holds the int value and the default value is nil
	var ptr *int
	fmt.Println(ptr)

	num := 10

	// Will get the memory address of the variable num
	refernce := &num
	fmt.Println(refernce)

	// Print the value stored in the memory address
	fmt.Println(*refernce)
}
