package main

import (
	"fmt"
)

func Prefix(a []int) {

	prefixSum := make([]int, len(a))
	prefixSum[0] = a[0]

	for i := 1; i < len(a); i++ {
		prefixSum[i] = prefixSum[i-1] + a[i]
	}

	fmt.Print(prefixSum)
}

func main() {
	a := []int{10, 20, 10, 5, 15}
	Prefix(a)
}
