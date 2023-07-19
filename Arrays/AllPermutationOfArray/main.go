package main

import "fmt"

func main() {

	nums := []int{0, 1}

	generateAllPermutation(nums, 0, len(nums)-1)
}

func generateAllPermutation(arr []int, left int, right int) {

	if left == right {
		fmt.Println(arr)

	} else {
		for i := left; i <= right; i++ {
			arr[left], arr[i] = arr[i], arr[left]
			generateAllPermutation(arr, left+1, right)
			arr[left], arr[i] = arr[i], arr[left]
		}
	}
}
