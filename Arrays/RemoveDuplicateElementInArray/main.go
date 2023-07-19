package main

import (
	"fmt"
)

// Work in case of sorted array only
func removeDuplicateFromSorted(arr []int) []int {

	result := []int{}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			result = append(result, arr[i])
		}
	}
	n := len(arr)
	result = append(result, arr[n-1])
	return result
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 3, 12, 5, 67, 9, 50, 11}
	fmt.Print(removeDuplicateFromSorted(arr))
}
