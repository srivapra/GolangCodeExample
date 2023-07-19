package main

import (
	"fmt"
	"math"
)

func findSecondLargestElement(arr []int) int {

	largest := math.MinInt
	secondLargest := math.MinInt

	arrayLength := len(arr)

	for i := 0; i < arrayLength; i++ {

		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			secondLargest = arr[i]
		}

	}
	return secondLargest

}

func main() {

	arr := []int{12, 8, 20, 43, 100, 65, 23, 76, 34}
	fmt.Println("The Second Largest Element in Array is : ", findSecondLargestElement(arr))

}
