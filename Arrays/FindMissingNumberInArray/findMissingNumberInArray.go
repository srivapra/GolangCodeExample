package main

import "fmt"

// Brute Force
/*
Create an temp array of size n+1 and fill with 0
Place 1 at temp array
Iterate through temp array and check if array is equal to 0 then return i
TC - O(2N)
SC - O(N)
*/
func findMissing(arr []int) int {
	temp := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		temp[i] = 0
	}
	fmt.Println(temp)

	for i := 0; i < len(arr); i++ {
		temp[arr[i]-1] = 1
	}
	for i := 1; i <= len(arr); i++ {
		if temp[i] == 0 {
			return i
		}

	}

	return -1
}

// Optimal - > TC - O(N) SC - O(1)
func findMissingNumberInArray(arr []int) int {

	n := len(arr)

	arraySum := 0

	sumOfNaturalNumber := ((n + 1) * (n + 2)) / 2

	for i := 0; i < n; i++ {
		arraySum = arraySum + arr[i]

	}

	return sumOfNaturalNumber - arraySum
}

func main() {

	arr := []int{1, 2, 4, 6, 3, 7, 8}
	fmt.Println(findMissingNumberInArray(arr))

}
