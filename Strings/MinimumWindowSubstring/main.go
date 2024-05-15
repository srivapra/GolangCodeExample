package main

import (
	"fmt"
	"sort"
)

/*
Generate All the substring of given string
for each substring check whether it contains all'
the given text/substring or not
find the smallest substring
*/
func generateSubstrings(s string) []string {
	substrings := []string{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substrings = append(substrings, s[i:j])
		}
	}
	return substrings
}

func containsPattern(substring, pattern string) bool {
	patternMap := make(map[rune]int)
	for _, char := range pattern {
		patternMap[char]++
	}

	for _, char := range substring {
		if val, ok := patternMap[char]; ok {
			patternMap[char]--
			if val == 1 {
				delete(patternMap, char)
			}
			if len(patternMap) == 0 {
				return true
			}
		}
	}
	return false
}

func findSmallestSubstring(stringVal, pattern string) string {
	substrings := generateSubstrings(stringVal)
	sort.Slice(substrings, func(i, j int) bool {
		return len(substrings[i]) < len(substrings[j])
	})

	for _, substring := range substrings {
		if containsPattern(substring, pattern) {
			return substring
		}
	}
	return ""
}

func main() {
	text := "ork"
	stringVal := "geeksforgeeks"

	smallestSubstring := findSmallestSubstring(stringVal, text)
	fmt.Println("Smallest substring containing pattern:", smallestSubstring)
}
