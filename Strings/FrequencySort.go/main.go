package main

import (
	"fmt"
	"sort"
	"strings"
)

func frequencySort(s string) string {
	m := make(map[byte][]byte)

	// groups characters by their frequencies.
	for i := 0; i < len(s); i++ {
		char := s[i]
		m[char] = append(m[char], char)
	}

	res := []string{}

	// iterates through the values of the map m (which are slices of characters grouped by frequency)
	// and appends them to the string slice res after converting them to strings.
	for _, val := range m {
		res = append(res, string(val))
	}

	// sort the string slices in descending order based on their lengths.
	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) > len(res[j])
	})

	return strings.Join(res, "")
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}
