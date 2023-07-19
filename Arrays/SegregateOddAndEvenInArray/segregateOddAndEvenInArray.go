package main

import "fmt"

func SegregateOddAndEvenInArray(arr []int) {

	left := 0
	right := len(arr) - 1

	for left < right {
		for arr[left]%2 == 0 && left < right {
			left++
		}
		for arr[right]%2 == 1 && left < right {
			right--
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

}

func main() {

	arr := []int{3, 21, 34, 12, 90, 32, 7, 1, 56}
	SegregateOddAndEvenInArray(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

}
