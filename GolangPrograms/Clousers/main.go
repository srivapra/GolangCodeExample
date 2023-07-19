// An annonymous function can make a clousers
// If you want to access the variable which is defined outside
// the function scope, It can be possible with the help of clousers

package main

import "fmt"

func main() {

	func() {

		fmt.Println("Annonymous Function")

	}()

	func(str string) {

		fmt.Println("Annonymous Function with arguments")

	}("Hi")

	value := func(val int) {

		fmt.Println("Assigned annonymous function into a variable")

	}

	value(5)

	num := 10
	fn := func() {

		fmt.Println("Value of num is : ", num)

	}

	fn()

}
