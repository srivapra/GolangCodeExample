package main

import "fmt"

func findDuplicateArrayElement(arr []int) {

	freqMap := make(map[int]int)

	// Iterate over the array element
	for _, value := range arr {

		// Increamenting the value of element by 1 in map
		freqMap[value] = freqMap[value] + 1
	}

	// Iterate over the frequency map and
	// print all the key which having values more than 1 (duplicate)
	for k, v := range freqMap {
		if v > 1 {
			fmt.Print(k, " ")

		}

	}

}

func main() {

	arr := []int{1, 1, 2, 2, 2, 3, 3, 3, 44, 4, 4}

	findDuplicateArrayElement(arr)

}
