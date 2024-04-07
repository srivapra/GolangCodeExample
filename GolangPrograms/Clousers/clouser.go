// In Golang, a closure is a function that references variables outside of its scope.
// A closure can outlive the scope in which it was created. Thus it can access variables within
// that scope, even after the scope is destroyed.

package main

import "fmt"

// Create a function that returns an anonymous function
func incrementor() func() int {

	i := 0

	return func() int {
		i++
		return i
	}
}

func incrementor1() func() int {

	return func() int {
		i := 0
		i++
		return i
	}
}

func main() {

	value := incrementor()
	fmt.Println(value())
	fmt.Println(value())
	fmt.Println(value())

	value1 := incrementor()
	fmt.Println(value1())

	value2 := incrementor1()
	fmt.Println(value2())
	fmt.Println(value2())
	fmt.Println(value2())

	value3 := incrementor1()
	fmt.Println(value3())

}
