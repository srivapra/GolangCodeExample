package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	result := append(nums1, nums2...)

	sort.Ints(result)

	n := len(result)

	if n%2 == 1 {
		return float64(result[n/2])
	} else {
		return float64(result[(n-1)/2]+result[n/2]) / 2
	}

	return -1

}

func main() {

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
