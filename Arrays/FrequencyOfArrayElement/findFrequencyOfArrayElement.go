package main

import "fmt"

func findFrequencyOfArrayElement(arr []int) map[int]int {

	freqMap := make(map[int]int)

	// Iterate over the array element
	for _, value := range arr {

		// Increamenting the value of element by 1 in map
		freqMap[value] = freqMap[value] + 1
	}

	return freqMap

}

func main() {

	arr := []int{12, 32, 17, 76, 45, 39, 55, 21}

	fmt.Println(findFrequencyOfArrayElement(arr))

}
