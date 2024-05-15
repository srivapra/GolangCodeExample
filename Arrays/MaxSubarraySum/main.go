package main

import (
	"fmt"
	"math"
)

/*
	Brute Force Approach

Generate all the subarray and calculate the sum of the subarray
and compare the subarray sum with maxSum and update the maxSum
*/

// Optimal
func maxSubarraySum(arr []int) int {
	sum := 0
	maxSum := math.MinInt
	arrLength := len(arr)

	for i := 0; i < arrLength; i++ {
		sum = sum + arr[i]
		if sum > maxSum {
			maxSum = sum

		}
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum

}

// If we carefully observe our algorithm, we can notice that the subarray
// always starts at the particular index where the sum variable is equal to 0, and at the ending index,
// the sum always crosses the previous maximum sum
func printSubarrayWithMaxSum(arr []int) int {

	sum := 0
	maxSum := math.MinInt
	arrLength := len(arr)

	ansStart := -1
	ansEnd := -1
	start := 0

	for i := 0; i < arrLength; i++ {
		if sum == 0 {
			start = i
		}

		sum = sum + arr[i]
		if sum > maxSum {
			maxSum = sum
			ansStart = start
			ansEnd = i
		}

		if sum < 0 {
			sum = 0
		}
	}

	for i := ansStart; i <= ansEnd; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	return maxSum

}

func main() {

	arr := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubarraySum(arr))

}
