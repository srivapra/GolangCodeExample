package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{2, 3, -2, 4}

	maxVal := arr[0]
	minVal := arr[0]
	maxProduct := arr[0]

	for i := 1; i < len(arr); i++ {

		// Swapping because, If I will multiply negative number in maxValue
		// then it will become negative number, After swapping If I will
		// multiply negative number in minValue then it will become positive number
		if arr[i] < 0 {
			minVal, maxVal = maxVal, minVal
		}

		minVal = int(math.Min(float64(arr[i]), float64(minVal*arr[i])))
		maxVal = int(math.Max(float64(arr[i]), float64(maxVal*arr[i])))

		if maxVal > maxProduct {
			maxProduct = maxVal
		}
	}

	fmt.Print(maxProduct)
}
