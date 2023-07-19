// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// // func isAlphaNumeric(str string) bool {
// // 	return regexp.MustCompile("^[a-zA-Z0-9]*$").MatchString(str)
// // }

// func main() {

// 	string := "Golang"

// 	isAlphaNumeric := isAlphaNumeric(string)
// 	if isAlphaNumeric {
// 		fmt.Printf("%s is an alpanumeric string  ", string)
// 	} else {
// 		fmt.Printf("%s is not an alpanumeric string  ", string)

// 	}
// }

// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func Example(s string) {
	fmt.Print("Hello")
}

func main() {
	go Example("Hi")
	fmt.Print("Anc")
	fmt.Print("Anc")
	fmt.Print("Anc")
	fmt.Print("Anc")
	fmt.Print("Anc")
	//time.Sleep(2 * time.Second)

}
