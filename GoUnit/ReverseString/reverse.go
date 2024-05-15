package main

import "fmt"

func Reverse(str string) string {
	rn := []rune(str)
	start := 0
	end := len(rn) - 1

	for start < end {
		rn[start], rn[end] = rn[end], rn[start]
		start++
		end--

	}
	return string(rn)
}

func main() {
	str := "prashant"
	fmt.Println(Reverse(str))
}
