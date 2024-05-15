package main

import (
	"fmt"
)

func PrintCharcterWithFreq(str string) {
	for i := 0; i < len(str); i++ {
		count := 1
		for i+1 < len(str) && str[i] == str[i+1] {
			i++
			count++
		}
		fmt.Printf("%c%d", str[i], count)
	}
}

func main() {
	str := "Javastudypoint"
	PrintCharcterWithFreq(str)
}
