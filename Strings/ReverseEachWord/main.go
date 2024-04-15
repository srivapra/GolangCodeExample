package main

import (
	"fmt"
	"strings"
)

func ReverseEachWord(str string) string {
	word := strings.Fields(str)

	for i, j := range word {
		word[i] = ReverseString(j)
	}
	return strings.Join(word, " ")
}

func ReverseString(str string) string {
	rn := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		rn[i], rn[j] = rn[j], rn[i]
	}
	return string(rn)
}

func main() {
	str := "Let's take LeetCode contest"
	fmt.Print(ReverseEachWord(str))
}
