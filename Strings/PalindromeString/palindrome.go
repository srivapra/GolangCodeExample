package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {

	rns := []rune(str)

	// Swapping the first character to last character
	// and last character to first and so on .......
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]

	}

	// If original string and reversed string is equal
	// then it is palindrome
	if string(rns) == str {
		return true
	}

	return false
}

func main() {

	string := "Level"

	// Converting the string to lower case
	string = strings.ToLower(string)

	fmt.Println("Input String : ", string)

	fmt.Println(isPalindrome(string))
}
