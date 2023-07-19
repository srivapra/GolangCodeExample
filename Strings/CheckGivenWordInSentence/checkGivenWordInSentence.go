package main

import (
	"fmt"
	"strings"
)

func checkGivenWordInSentence(sentence string, word string) bool {

	sentenceArray := strings.Fields(sentence)

	// Iterate over the string array and check
	// if the word in the sentence array is same
	// as word
	for i := 0; i < len(sentenceArray); i++ {
		if sentenceArray[i] == word {
			fmt.Println(sentenceArray[i])
			return true

		}
	}

	return false
}

func main() {

	sentence := "GFG is a computer science portal"

	sentence = strings.ToLower(sentence)

	word := "Computer"

	word = strings.ToLower(word)

	fmt.Println(checkGivenWordInSentence(sentence, word))
}
