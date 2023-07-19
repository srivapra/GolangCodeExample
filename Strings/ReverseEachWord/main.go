package main

import (
	"fmt"
	"strings"
)

func ReverseEachWord(s string) {
	splitString := strings.Split(s, " ")

	for _, v := range splitString {
		rn := []rune(v)
		for i, j := 0, len(rn)-1; i < j; i, j = i+1, j-1 {
			rn[i], rn[j] = rn[j], rn[i]
		}
		fmt.Println(string(rn))
	}

}

func main() {
	str := "Hi Hello World"
	ReverseEachWord(str)
}
