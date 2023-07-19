package main

import "fmt"

func findMinMaxElementInArray(arr []int) {

	// Assigning first element of the array
	// as minimum and Maximum
	min := arr[0]
	max := arr[0]

	// Iterate over the array and check if the array element
	// is smaller than min, then updating the value of min, else
	// if the array element is greater than max,
	// then updating the value of max
	for _, value := range arr {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}

	fmt.Println("Min : ", min, "\nMax : ", max)
}

func main() {

	arr := []int{10, 21, 2, 98, 10016, 32, 65, 100}

	findMinMaxElementInArray(arr)

}
