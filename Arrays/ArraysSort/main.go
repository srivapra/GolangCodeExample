package main

import (
	"fmt"
)

func SortArray(arr1 []int, arr2 []int) []int {
	left := 0
	right := len(arr2) - 1
	var res []int

	for left < len(arr1) && right >= 0 {
		if arr1[left] < arr2[right] {
			res = append(res, arr1[left])
			left++

		} else {
			res = append(res, arr2[right])
			right--

		}
		fmt.Println(left, right)
		//fmt.Println(right)
	}
	for right >= 0 {
		res = append(res, arr2[right])
		right--

	}
	for left < len(arr1) {
		res = append(res, arr1[left])
		left++

	}
	return res

}

func main() {
	arr1 := []int{0, 2, 4, 6}
	arr2 := []int{7, 5, 2, 1}
	fmt.Print(SortArray(arr1, arr2))
}
