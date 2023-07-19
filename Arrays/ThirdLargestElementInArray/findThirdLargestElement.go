package main

import (
	"fmt"
	"math"
)

func findThirdLargestElement(arr []int) int {

	largest := math.MinInt
	secondLargest := math.MinInt
	thirdLargest := math.MinInt

	arrayLength := len(arr)

	for i := 0; i < arrayLength; i++ {

		if arr[i] > largest {
			thirdLargest = secondLargest
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			thirdLargest = secondLargest
			secondLargest = arr[i]
		} else if arr[i] > thirdLargest && arr[i] != secondLargest {
			thirdLargest = arr[i]
		}

	}
	return thirdLargest

}

func main() {

	arr := []int{12, 8, 20, 43, 100, 65, 23, 76, 34}
	fmt.Println("The Third Largest Element in Array is : ", findThirdLargestElement(arr))

}
