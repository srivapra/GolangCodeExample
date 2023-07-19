package main

import "fmt"

func findLargestElementInArray(arr []int) int {

	// Assigning first element of the array as maximum
	max := arr[0]

	// Iterate over the array and check the array element
	// is greater than max, then updating the value of max
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	return max
}

func main() {

	arr := []int{10, 21, 2, 98, 10016, 32, 65, 100}

	fmt.Println(findLargestElementInArray(arr))

}
