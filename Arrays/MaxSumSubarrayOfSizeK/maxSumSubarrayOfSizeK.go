package main

import (
	"fmt"
)

func MaxSumSubarrayOfSizeK(arr []int, k int) int {

	wSum := 0
	mSum := 0
	arrLength := len(arr)

	if arrLength < k {
		fmt.Println("Array Size Must be greater than the value of K")
		return -1
	}

	// Calculating the sum of the window of size k
	// i.e. Calculating the sum of first k array element
	for i := 0; i < k; i++ {
		wSum = wSum + arr[i]
	}

	// Ignore the first k element and start iterating the
	// array element from K
	for end := k; end < arrLength; end++ {

		// New Window sum is the addition of a new
		// element arr[end] and removal of the previous element arr[end-k]
		wSum += arr[end] - arr[end-k]

		// Finding the max values amoung wSum and mSum
		if wSum > mSum {
			mSum = wSum
		}

	}

	return mSum
}

func main() {

	arr := []int{1, 9, -1, -2, 7, 3, -1, 2}
	k := 9
	fmt.Println(MaxSumSubarrayOfSizeK(arr, k))

}
