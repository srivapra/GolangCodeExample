package main

import "fmt"

func linearSearchInArray(arr []int, item int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return i
		}

	}
	return -1
}

func main() {

	arr := []int{12, 56, 43, 21, 89, 100, 43}
	item := 12

	fmt.Println(linearSearchInArray(arr, item))

}
