package main

import "fmt"

func Segregate0And1InArray(arr []int) {

	left := 0
	right := len(arr) - 1

	for left < right {
		for arr[left] == 0 && left < right {
			left++
		}
		for arr[right] == 1 && left < right {
			right--
		}
		if left < right {
			arr[left] = 0
			arr[right] = 1
			left++
			right--
		}
	}

}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func main() {

	arr := []int{0, 1, 1, 0, 0, 1, 0, 0, 1}
	Segregate0And1InArray(arr)
	printArray(arr)

}
