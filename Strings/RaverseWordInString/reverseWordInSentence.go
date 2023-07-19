package main

import (
	"fmt"
	"strings"
)

func reverseWordInSentence(str string) string {

	// Split the given string
	trimedStr := strings.Trim(str, " ")
	words := strings.Fields(trimedStr)

	// Swap the word of the sentence like
	// first with last, and so on ..
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {

		words[i], words[j] = words[j], words[i]

	}

	// Joining the string by seprator " "
	return strings.Join(words, " ")
}

func main() {

	sentence1 := "the sky is blue"
	sentence2 := "  hello world  "
	sentence3 := "a good   example"

	fmt.Println(reverseWordInSentence(sentence1))
	fmt.Println(reverseWordInSentence(sentence2))
	fmt.Println(reverseWordInSentence(sentence3))
}
