package main

import "fmt"

func frequencyMap(arr string) map[string]int {

	// Creating a frequency map to store the frequency
	// of characters
	freq := make(map[string]int)

	// Iterate over the string character by character
	for _, char := range arr {

		// Increamenting the value of character by 1
		freq[string(char)] = freq[string(char)] + 1
	}

	return freq
}

func main() {

	arr := "prashant"

	freq_map := frequencyMap(arr)

	fmt.Println("The frequency array is : ", freq_map)

}
