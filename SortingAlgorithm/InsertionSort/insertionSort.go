package main

import "fmt"

func insertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current
	}
}

func main() {

	arr := []int{12, 6, 78, 32, 10, 9, 11}
	n := len(arr)

	insertionSort(arr)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

}
