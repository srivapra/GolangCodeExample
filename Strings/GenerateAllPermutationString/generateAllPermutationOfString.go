package main

import "fmt"

func generateAllPermutation(sampleRune []rune, left, right int) {

	if left == right {
		fmt.Println(string(sampleRune))
	} else {
		for i := left; i <= right; i++ {

			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]

			generateAllPermutation(sampleRune, left+1, right)

			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
		}
	}
}

func main() {

	str := "abc"

	sampleRune := []rune(str)

	generateAllPermutation(sampleRune, 0, len(sampleRune)-1)

}
