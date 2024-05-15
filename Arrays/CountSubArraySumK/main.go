package main

import (
	"fmt"
)

// Generate all the subarray and check how many
// subarrays having sum equal to k then return count
// TC - O(N3)
func countSubArraySumBrute(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]

			}
			if sum == k {
				count++
			}
		}
	}
	return count
}

// If we carefully observe, we can notice that to get the sum of the current subarray
// we just need to add the current element(i.e. arr[j]) to the sum of the previous subarray.
// we can remove the third loop and while moving j pointer, we can calculate the sum
// TC - O(N2)
func countSubArraySumBetter(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// TC - O(N)
// SC - O(N)
func countSubarraySumOptimal(nums []int, k int) int {
	count := 0
	prefixSum := 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {

		// the prefix sum of a subarray ending at index i
		prefixSum += nums[i]

		// Check if (prefixSum - k) is there is the map or not
		// if it is there then incremenet the count by 1. If it is not
		// there simply put the element in map increasing its occurrence by 1
		count += m[prefixSum-k]
		m[prefixSum]++
	}
	return count
}

func main() {
	arr := []int{1, 1, 1}
	k := 2
	fmt.Println(countSubarraySumOptimal(arr, k))

}
