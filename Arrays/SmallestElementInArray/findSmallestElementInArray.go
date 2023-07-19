package main

import "fmt"

func findSmallestElementInArray(arr []int) int {

	// Assigning first element of the array as minimum
	min := arr[0]

	// Iterate over the array and check the array element
	// is smaller than min, then updating the value of min
	for _, value := range arr {
		if value < min {
			min = value
		}
	}

	return min
}

func main() {

	arr := []int{10, 21, 2, 98, 10016, 32, 65, 100}

	fmt.Println(findSmallestElementInArray(arr))

}
