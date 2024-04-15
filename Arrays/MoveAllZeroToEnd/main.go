package main

import "fmt"

/*
1. Place all the non zero element in temp array -> O(N)
2. Place the element of temp array in original array -> 	O(X) - x is the element in temp array
3. Place all the zeros at end -> O(N-X)
TC = O(N) + O(X) + O(N-X) -> O(2N)
SC = O(N)
*/
func MoveZeroNaive(arr []int) {
	var temp []int

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			temp = append(temp, arr[i])
		}
	}

	for i := 0; i < len(temp); i++ {
		arr[i] = temp[i]
	}

	for i := len(temp); i < len(arr); i++ {
		arr[i] = 0
	}
	fmt.Println(arr)
}

func moveZeroesOptimal(nums []int) {

	j := -1

	// Find the index index of first zero element
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j = i
			break
		}
	}

	// Iterate the array and check if arr[i] != 0 then swap it and increment j
	for i := j + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	fmt.Println(nums)

}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	MoveZeroNaive(arr)
}
