package main

import "fmt"

func binarySearch(arr []int, item int, si int, ei int) int {

	for si <= ei {

		mid := si + (ei-si)/2

		if item == arr[mid] {
			return mid

		}

		if item < arr[mid] {
			ei = mid - 1
		} else {
			si = mid + 1
		}

	}
	return -1
}

func main() {

	arr := []int{12, 56, 78, 89, 90, 95}

	n := len(arr)

	item := 90

	fmt.Println(binarySearch(arr, item, 0, n-1))

}
