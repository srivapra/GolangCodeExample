package main

import "sort"

func threeSum(arr []int) [][]int {

	sort.Ints(arr)

	var result [][]int

	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left := i + 1
		right := len(arr) - 1

		for left < right {
			total := arr[i] + arr[left] + arr[right]
			if total < 0 {
				left++
			} else if total > 0 {
				right--

			} else {
				result = append(result, []int{arr[i], arr[left], arr[right]})
			}

			for left < right && arr[left] == arr[left+1] {
				left++
			}
			for left < right && arr[right] == arr[right-1] {
				right--
			}
			left++
			right--

		}

	}

	return result
}

func main() {

	arr := []int{-1, 0, 1, 2, -1, -4}

	sort.Ints(arr)

	for i := 1; i < len(arr)-2; i++ {

	}
}
