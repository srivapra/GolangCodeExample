package main

import (
	"fmt"
	"sort"
)

func threeSum(arr []int, sum int) [][]int {

	sort.Ints(arr)

	var result [][]int

	for i := 0; i < len(arr)-2; i++ {
		left := i + 1
		right := len(arr) - 1
		for left < right {
			if arr[i]+arr[left]+arr[right] == sum {
				result = append(result, []int{arr[i], arr[left], arr[right]})
				left++
				right--
			} else if arr[i]+arr[left]+arr[right] < sum {
				left++
			} else {
				right--
			}
		}
	}
	return result

}

func main() {

	arr := []int{1, 4, 45, 6, 10, 8}
	sum := 22

	fmt.Println(threeSum(arr, sum))

}
