package main

import (
	"fmt"
	"sort"
)

func maxOccuringCharacter(arr string) string {

	// Creating a frequency map to store the frequency
	// of characters
	freq := make(map[string]int)

	// Iterate over the string character by character
	for _, value := range arr {

		// Increamenting the value of character by 1
		freq[string(value)] = freq[string(value)] + 1
	}

	// Creating a slice
	keys := make([]string, 0, len(freq))

	// Iterate over the map and
	// appending the key of the map
	// to slice (keys)
	for key := range freq {
		keys = append(keys, key)
	}

	fmt.Println("Frequency Map Before Sorting : ", freq)
	fmt.Println("Finding Only Keys from the map : ", keys)

	// Soring the map by values .. In descending order
	sort.SliceStable(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	fmt.Println("Keys After Sorting based on values: ", keys)

	// Returning first element of the keys
	return keys[0]

}
func main() {

	str := "javastudypoint"
	fmt.Println("Most Occuring Character is : ", maxOccuringCharacter(str))

}
