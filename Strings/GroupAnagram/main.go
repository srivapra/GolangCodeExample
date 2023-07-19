package main

import (
	"fmt"
	"sort"
)

// Iterate over the strings and group them based on the frequency of the
// characters in the string.
// func groupAnagrams(strs []string) [][]string {
// 	m := make(map[[26]int][]string)
// 	for _, s := range strs {
// 		var freq [26]int
// 		for _, c := range s {
// 			freq[c-'a']++
// 		}
// 		m[freq] = append(m[freq], s)
// 	}

// 	var groups [][]string
// 	for _, v := range m {
// 		groups = append(groups, v)
// 	}
// 	return groups
// }

func groupAnagrams(words []string) [][]string {
	anagramGroups := make(map[string][]string)
	for _, word := range words {
		sortedWord := sortString(word)
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}
	result := make([][]string, 0, len(anagramGroups))
	for _, group := range anagramGroups {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Print(groupAnagrams(str))
}
