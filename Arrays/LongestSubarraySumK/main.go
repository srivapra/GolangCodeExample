package main

import "fmt"

/*
Generate all the subarray and calculate the sum of the subarray and check if the
sum is equal to k then update the maxlength and return
TC := O(N2)
SC := O(1)
*/
func LongestSubarrayWithSumK(arr []int, k int) int {

	maxLength := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum == k {
				maxLength = Max(maxLength, j-i+1)
				//fmt.Println(arr[i : j+1])
			}
		}
	}
	return maxLength
}

/*
Find sum of subarray
check if sum < k then move ahead means j++
check if sum == k, then update the maxlenth
check if sum > k, then shrink the previous element by moving i to ahead meeans i++ until the sum > k
decrement the sum also
After the while loop increment j++

TC := O{N}
SC := O{1}
*/
func FindLongestSubarraySumK(arr []int, k int) int {
	i, j := 0, 0
	sum := 0
	maxLength := 0

	for j < len(arr) {
		sum += arr[j]
		if sum < k {
			j++
		} else if sum == k {
			maxLength = Max(maxLength, j-i+1)
			j++
		} else if sum > k {
			for sum > k {
				sum -= arr[i]
				i++
			}
			j++
		}
	}
	return maxLength

}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	arr := []int{1, 2, 3, 1, 1, 1, 1, 4, 2, 3}
	k := 3
	fmt.Println(FindLongestSubarraySumK(arr, k))
}
