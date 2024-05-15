package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	ans := Add(4, 7)
	fmt.Println(ans)
}
