package main

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	m := make(map[byte][]byte)

	for i := 0; i < len(s); i++ {
		char := s[i]
		m[char] = append(m[char], char)
	}

	res := []string{}

	for _, val := range m {
		res = append(res, string(val))
	}

	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) > len(res[j])
	})

	return strings.Join(res, "")
}

func main() {
	s := "tree"
	frequencySort(s)
}
