package main

import "fmt"

func longestSubStringWithoutRepeatingChar(str string) int {

	// Create a map to store characters and their indices
	m := make(map[byte]int)
	start := 0
	maxLength := 0
	longestSubstring := ""

	for i := 0; i < len(str); i++ {
		// Check if the character is already in the map and its index is after the start index
		if index, ok := m[str[i]]; ok && index >= start {
			// Move the start index to the next index after the repeated character
			start = index + 1
		}
		// Update the index of the character
		m[str[i]] = i

		// Update the maximum length
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			longestSubstring = str[start : i+1]
		}

	}
	fmt.Println(longestSubstring)
	return maxLength

}

func main() {
	str := "pwwkew"
	fmt.Println(longestSubStringWithoutRepeatingChar(str))
}
