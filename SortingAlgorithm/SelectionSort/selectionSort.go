package main

import "fmt"

func selectionSort(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {

		// Find the index of the minimum element
		// from unsorted array
		midIndex := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[midIndex] {
				midIndex = j

			}
		}

		// Swap the minimum element with first element of unsorted array
		arr[midIndex], arr[i] = arr[i], arr[midIndex]
	}

}

func main() {

	arr := []int{12, 6, 78, 32, 10, 9, 11}
	n := len(arr)

	selectionSort(arr)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

}
