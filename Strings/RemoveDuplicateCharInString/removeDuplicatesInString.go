package main

import (
	"fmt"
)

func removeDuplicateCharacters(str string) string {

	// convert a string into array of characters
	// (which is slice of runes) having ascci value of character
	char := []rune(str)

	// Creating a bool map (empty map)
	inResult := make(map[string]bool)

	result := ""

	// Iterate over the array of characters
	for _, str := range char {

		// Coverting the rune into string
		char := string(str)

		// Checking char is present inside the map
		// or not, If not then appending the char into result variable
		// and marking this character as true in map
		if _, ok := inResult[char]; !ok {
			inResult[char] = true
			result = result + char
		}

	}
	return result
}

func main() {

	str := "aabb"

	fmt.Println(removeDuplicateCharacters(str))
}
