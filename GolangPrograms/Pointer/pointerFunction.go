package main

import "fmt"

func pointerFunction(str *string) {

	fmt.Println(*str)

}

func main() {

	str := "Hello"
	pointerFunction(&str)
}
