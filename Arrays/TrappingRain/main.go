package main

import (
	"fmt"
)

/*
Traverse every array element and find the highest bars on the left and right sides.
Take the smaller of two heights. The difference between the smaller height and the height
of the current element is the amount of water that can be stored in this array element.
*/

func MaxWater(arr []int) int {
	res := 0

	for i := 1; i < len(arr)-1; i++ {

		// Find maximum element on its left
		left := arr[i]
		for j := 0; j < i; j++ {
			left = Max(left, arr[j])
		}

		// Find maximum element on its right
		right := arr[i]
		for j := i + 1; j < len(arr); j++ {
			right = Max(right, arr[j])
		}

		// Update maximum water value
		res += Min(left, right) - arr[i]

	}
	return res

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
For every element we can precalculate and store the highest bar on the left and
on the right (say stored in arrays left[] and right[]).
Then iterate the array and use the precalculated values to find the amount of water
stored in this index,
which is the same as ( min(left[i], right[i]) – arr[i] )
*/
// O(N) 0(N)
func findWater(arr []int) int {
	n := len(arr)

	// left[i] contains height of tallest bar to the
	// left of i'th bar including itself
	left := make([]int, n)

	// Right[i] contains height of tallest bar to
	// the right of ith bar including itself
	right := make([]int, n)

	// Initialize result
	water := 0

	// Fill left array
	left[0] = arr[0]
	for i := 1; i < n; i++ {
		left[i] = Max(left[i-1], arr[i])
	}

	// Fill right array
	right[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = Max(right[i+1], arr[i])
	}

	// Calculate the accumulated water element by
	// element consider the amount of water on i'th bar,
	// the amount of water accumulated on this
	// particular bar will be equal to min(left[i],
	// right[i]) - arr[i] .
	for i := 0; i < n; i++ {
		water += Min(left[i], right[i]) - arr[i]
	}

	return water
}

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(MaxWater(arr))
}
