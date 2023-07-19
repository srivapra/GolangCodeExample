package main

import "fmt"

func bubbleSort(arr []int) {

	arrayLength := len(arr)

	for i := 0; i < arrayLength-1; i++ {
		for j := 0; j < arrayLength-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}

}

func main() {

	arr := []int{12, 65, 32, 10, 8, 56, 43}

	bubbleSort(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

}
