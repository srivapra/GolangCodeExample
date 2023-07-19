package main

import "fmt"

func searchRange(nums []int, target int) []int {

	result := make([]int, 2)
	result[0] = findStartingIndex(nums, target)
	result[1] = findEndingIndex(nums, target)
	return result

}

func findStartingIndex(nums []int, target int) int {
	index := -1

	start := 0
	end := len(nums) - 1

	for start <= end {
		midPoint := start + (end-start)/2

		if nums[midPoint] >= target {
			end = midPoint - 1
		} else {
			start = midPoint + 1
		}
		if nums[midPoint] == target {
			index = midPoint
		}
	}
	return index

}

func findEndingIndex(nums []int, target int) int {
	index := -1

	start := 0
	end := len(nums) - 1

	for start <= end {
		midPoint := start + (end-start)/2

		if nums[midPoint] <= target {
			start = midPoint + 1
		} else {
			end = midPoint - 1
		}

		if nums[midPoint] == target {
			index = midPoint
		}
	}
	return index

}

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	target := 8

	fmt.Print(searchRange(arr, target))
}
