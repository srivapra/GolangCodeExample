package main

import "fmt"

func duplicateCharacters(arr string) {

	// Creating a frequency map to store the frequency
	// of characters
	freq := make(map[string]int)

	// Iterate over the string character by character
	for _, char := range arr {

		// Increamenting the value of character by 1
		freq[string(char)] = freq[string(char)] + 1
	}

	fmt.Println("Frequency Map : ", freq)

	// Iterate over the frequency map and
	// print all the key which having values more than 1 (duplicate)
	for k, v := range freq {
		if v > 1 {
			fmt.Println("Duplicate Characters are : ", k)
		}
	}

}

func main() {

	string := "geeksforgeeks"

	duplicateCharacters(string)

}
