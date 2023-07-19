package main

import "fmt"

func maxSubarraySum(arr []int) int {

	sum := 0
	maxSum := arr[0]
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

func main() {

	arr := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubarraySum(arr))

}
