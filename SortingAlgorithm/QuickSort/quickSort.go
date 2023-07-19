package main

import "fmt"

func partition(arr []int, low int, high int) int {

	// Choosing pivot as last element
	pivot := arr[high]

	// index of smaller element and indicates the right position
	// of the pivot found so far
	i := (low - 1)

	for j := low; j <= high-1; j++ {

		// If current element is smaller
		// than the pivot
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]

		}
	}

	// Place pivot at its correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// Return index of pivot
	return (i + 1)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {

	arr := []int{12, 6, 78, 32, 10, 9, 11}
	n := len(arr)
	quickSort(arr, 0, n-1)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

}
