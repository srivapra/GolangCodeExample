package main

import "fmt"

func maxOccuringCharacterWithoutSort(arr string) string {

	// Creating a frequency map to store the frequency
	// of characters
	freq := make(map[string]int)

	// Iterate over the string character by character
	for _, char := range arr {

		// Increamenting the value of character by 1
		freq[string(char)] = freq[string(char)] + 1
	}

	fmt.Println("Frequency Map : ", freq)

	maxK := ""
	maxV := 0

	// Iterate over the frequency map and
	// update the maxV and maxK and return the maxK
	for k, v := range freq {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}

	// maxK is the key with maximum value
	return maxK
}

func main() {

	string := "geeksforgeeks"

	most_occuring := maxOccuringCharacterWithoutSort(string)

	fmt.Println("Most Occuring Character is : ", most_occuring)

}
