package main

import (
	"fmt"
)

func checkTwoStringAnagram(str1 string, str2 string) bool {

	len1 := len(str1)
	len2 := len(str2)

	if len1 != len2 {
		return false
	}

	// Creating a frequency map to store the frequency
	// of characters
	freq1 := make(map[string]int)

	//freq2 := make(map[string]int)

	// Iterate over the string character by character
	for i := 0; i < len1 && i < len2; i++ {

		// Increamenting the value of character by 1
		freq1[string(str1[i])] = freq1[string(str1[i])] + 1

		freq1[string(str2[i])] = freq1[string(str2[i])] - 1
	}

	for _, v := range freq1 {
		if v != 0 {
			return false
		}
	}

	// // Comparing both frequencies
	// result := reflect.DeepEqual(freq1, freq2)

	// if !result {
	// 	return false
	// }

	return true

}

func main() {

	str1 := "study"
	str2 := "dusty"

	fmt.Println(checkTwoStringAnagram(str1, str2))

}
