// panic means an unexpected condition arises in your Go program due to which
// the execution of your program is terminated. Sometimes panic occurs at runtime
// when some specific situation arises like out-of-bounds array accesses,
// integer divide by zero etc.

package main

import "fmt"

func main() {

	a := 10
	b := 0
	c := a / b

	fmt.Print(c) // throw panic

}
