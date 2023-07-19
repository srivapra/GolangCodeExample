package main

import (
	"fmt"
	"strings"
)

func checkIfPangram(sentence string) bool {
	for i := int('a'); i <= int('z'); i++ {
		if !strings.Contains(sentence, string(i)) {
			return false
		}
	}
	return true
}

func main() {

	str := "abcdefghijklmnopqrstuvwxyz"

	lowercase := strings.ToLower(str)

	fmt.Println(checkIfPangram(lowercase))

}
